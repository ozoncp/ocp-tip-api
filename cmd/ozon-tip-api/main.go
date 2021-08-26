package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-tip-api/internal/api"
	configuration "github.com/ozoncp/ocp-tip-api/internal/config"
	"github.com/ozoncp/ocp-tip-api/internal/db"
	"github.com/ozoncp/ocp-tip-api/internal/metrics"
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
	httpPort           = ":8081"
	metricsPort        = ":9100"
)

func run(dbConn *sqlx.DB, producer sarama.SyncProducer) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpTipApiServer(s, api.NewOcpTipApi(repo.NewRepo(dbConn), producer))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func newProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpTipApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		panic(err)
	}
}

func runMetrics() {
	metrics.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(metricsPort, nil)
	if err != nil {
		panic(err)
	}
}

func initTracing(jaegerHost string) {
	config := jaegercfg.Configuration{
		ServiceName: "ocp_tip_api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHost,
		},
	}

	logger := jaegerlog.StdLogger
	metricsFactory := jaegermetrics.NullFactory

	tracer, closer, err := config.NewTracer(jaegercfg.Logger(logger), jaegercfg.Metrics(metricsFactory))

	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)
	_ = closer.Close()
}

func main() {
	go runJSON()
	go runMetrics()

	config, err := configuration.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)
	dbConn, err := db.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	producer, err := newProducer(config.Brokers)
	if err != nil {
		log.Fatal(err)
	}

	initTracing(config.JaegerAgentHostPort)
	if err := run(dbConn, producer); err != nil {
		log.Fatal(err)
	}
}
