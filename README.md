# goGrpc
go grpc example include set matedata and grpc-gateway

How to do install GRPC

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go get -u github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
```

How to do generate protoc

```bash
protoc --proto_path=./grpc --go_out=./grpc --plugin=protoc-gen-go=../bin/protoc-gen-go myGrpc.proto
protoc --proto_path=./grpc --go-grpc_out=./grpc --plugin=protoc-gen-go-grpc=../bin/protoc-gen-go-grpc myGrpc.proto
protoc ./grpc --grpc-gateway_out=./grpc --plugin=protoc-gen-grpc-gateway=../bin/protoc-gen-grpc-gateway myGrpc.proto
```

How to do run

```bash
go run server.go
go run client.go
```

How use make file

```bash
make all
```

How to do HTTP CALL

```bash
curl -X GET -k http://127.0.0.1:8080/v1/md5?inString=world
curl -X POST -k http://127.0.0.1:8080/v1/mathSum -d '{"min": 4, "max": 5}'
curl -X GET -H "Grpc-Metadata-appid:zhang" -H "Grpc-Metadata-appkey:woods" -k http://127.0.0.1:8080/v1/md5?inString=world
```
