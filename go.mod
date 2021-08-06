module grpc

require (
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jsonrpc2 v0.0.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
)

replace github.com/jsonrpc2 => ../github.com/jsonrpc2

go 1.15
