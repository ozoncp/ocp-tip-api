package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tip-api/internal/api"
	configuration "github.com/ozoncp/ocp-tip-api/internal/config"
	"github.com/ozoncp/ocp-tip-api/internal/db"
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
	httpPort           = ":8081"
)

func run(dbConn *sqlx.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpTipApiServer(s, api.NewOcpTipApi(repo.NewRepo(dbConn)))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
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

func main() {
	go runJSON()

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

	if err := run(dbConn); err != nil {
		log.Fatal(err)
	}
}
