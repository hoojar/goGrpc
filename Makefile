TARGET=grpc
outPath=./
binPath=${GOPATH}/bin/
protoPath=-I ./ -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
all: clean build

clean:
	rm -rf $(TARGET)-client
	rm -rf $(TARGET)-server

build:
	go build -o jsonrpc2 jsonrpc2.go
	go build -o ${TARGET}-client client.go
	go build -o $(TARGET)-server server.go

proto:
	protoc ${protoPath} --go_out=$(outPath) --plugin=protoc-gen-go=${binPath}protoc-gen-go myGrpc.proto
	protoc ${protoPath} --go-grpc_out=$(outPath) --plugin=protoc-gen-go-grpc=${binPath}protoc-gen-go-grpc myGrpc.proto
	protoc ${protoPath} --grpc-gateway_out=$(outPath) --plugin=protoc-gen-grpc-gateway=${binPath}protoc-gen-grpc-gateway myGrpc.proto

