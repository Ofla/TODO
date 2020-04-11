GOOGLE_APIS=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

install-tools:
	go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get  github.com/golang/protobuf/protoc-gen-go


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

test:
	go ./buslog/todo_test.go