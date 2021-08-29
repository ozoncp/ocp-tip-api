package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	"os"
	"os/signal"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
	httpPort           = ":8081"
	metricsPort        = ":9100"
)

func newProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func newGrpcGatewayServer(ctx context.Context) (*http.Server, error) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpTipApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}, nil
}

func newMetricsServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    metricsPort,
		Handler: mux,
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	metricsServer := newMetricsServer()
	go func() {
		metrics.RegisterMetrics()
		if err := metricsServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	grpcGatewayServer, err := newGrpcGatewayServer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := grpcGatewayServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	producer, err := newProducer(config.Brokers)
	if err != nil {
		log.Fatal(err)
	}

	initTracing(config.JaegerAgentHostPort)

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	desc.RegisterOcpTipApiServer(grpcServer, api.NewOcpTipApi(repo.NewRepo(dbConn), producer))
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop

	if err := metricsServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	if err := grpcGatewayServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	grpcServer.GracefulStop()
}
