module github.com/ozoncp/ocp-tip-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Shopify/sarama v1.29.1
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40 // indirect
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.10.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.23.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210819135213-f52c844e1c1c // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210820002220-43fce44e7af1 // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api => ./pkg/ocp-tip-api
