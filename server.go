package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	gi "grpc/api"
	pb "grpc/myGrpc"
)

func main() {
	var restAddr = flag.String("restAddr", ":8080", "bind rest address") //REST绑定的IP:端口号
	var grpcAddr = flag.String("grpcAddr", ":8081", "bind grpc address") //GRPC绑定的IP:端口号
	flag.Parse()                                                         //解析命令行参数

	go startGrpc(*grpcAddr)
	startRest(*restAddr, *grpcAddr)
}

func startGrpc(grpcAddr string) {
	lis, err := net.Listen("tcp", grpcAddr) //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	var rpc = grpc.NewServer()                   //创建gRPC服务
	pb.RegisterProdServiceServer(rpc, gi.RpcApi) //注册接口服务
	log.Printf("TCP GRPC API [%s] start\n", grpcAddr)
	if err = rpc.Serve(lis); err != nil { // 将监听交给gRPC服务处理
		log.Fatalf("failed to serve: %v", err)
	}
}

func startRest(restAddr string, grpcAddr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterProdServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("HTTP REST API [%s] start", restAddr)
	if err = http.ListenAndServe(restAddr, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
