RELEASE?=$(shell git describe --tags --candidates=1 | grep -P '\d+\.\d+\.\d+' -o)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-tip-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-tip-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-tip-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-tip-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-tip-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-tip-api/ocp-tip-api.proto
		mv pkg/ocp-tip-api/github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api/* pkg/ocp-tip-api/
		rm -rf pkg/ocp-tip-api/github.com
		mkdir -p cmd/ozon-tip-api

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build \
                -ldflags "-s -w -X github.com/ozoncp/ocp-tip-api/internal/version.Release=${RELEASE} \
                -X github.com/ozoncp/ocp-tip-api/internal/version.Commit=${COMMIT} \
                -X github.com/ozoncp/ocp-tip-api/internal/version.BuildTime=${BUILD_TIME}" \
                -o bin/ocp-tip-api cmd/ozon-tip-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-tip-api
		cp api/ocp-tip-api/ocp-tip-api.proto vendor.protogen/api/ocp-tip-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate

run:
	go run cmd/ozon-tip-api/main.go