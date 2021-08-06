package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "grpc/myGrpc"
)

func main() {
	/*
		建立连接到gRPC服务
	*/
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() //函数结束时关闭连接

	/*
		模拟请求数据 os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	*/
	var str string = "test123"
	if len(os.Args) > 1 {
		str = os.Args[1]
	}

	/*
		调用gRPC接口
	*/
	rpcClient := pb.NewProdServiceClient(conn) //创建服务的客户端
	md := metadata.Pairs("appid", "zhang", "appkey", "woods")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	if res, err := rpcClient.Md5(ctx, &pb.Req{InString: str}); err != nil {
		log.Fatalf("%v", err)
	} else {
		log.Printf("%s Md5 is: %s", str, res.RetString)
	}

	mathRes, err := rpcClient.MathSum(context.Background(), &pb.NumRequest{Min: 1, Max: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Math sum amount: %d", mathRes.Amount)
}
