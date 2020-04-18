VERSION=1.0.1

GOOGLE_APIS=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

install-tools:
	$(GOGET)  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET)  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET)  github.com/golang/protobuf/protoc-gen-go


generate-all: install-tools

	@rm -f www/swagger.json

	@protoc \
		-I. \
		-I$(GOPATH)/src/$(GOOGLE_APIS) \
		--go_out=plugins=grpc:. \
		--swagger_out=logtostderr=true:. \
		--grpc-gateway_out=logtostderr=true:. \
		proto/v1/todo.proto

	@cp proto/v1/todo.swagger.json www/swagger.json

runGRPCServer:
	go run ./server/grpcServer.go

runHTTPServer:
	go run ./http/httpServer.go

runGRPCClient:
	go run ./client/grpcClient.go

init:
	$(GOCMD) mod download

test:
	$(GOTEST) -v ./...