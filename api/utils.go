package api

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "grpc/myGrpc"
)

// 业务实现方法的容器
type rpcService struct {
	pb.UnimplementedProdServiceServer
}

//var RpcApi = rpcService{}
var RpcApi = new(rpcService)

// 为server定义 Md5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *pb.Req[相应接口定义的请求参数])
// 返回 (*pb.Res[相应接口定义的返回参数，必须用指针], error)
func (s *rpcService) Md5(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "无Token认证信息")
	}

	jsonStr, _ := json.Marshal(md)
	log.Println(string(jsonStr))

	var appid, appkey string
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "zhang" || appkey != "woods" {
		return nil, status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	log.Printf("Received: %v Token info: appid=%s,appkey=%s", in.GetInString(), appid, appkey)
	fmt.Println("MD5方法请求JSON:" + in.InString)
	return &pb.Res{RetString: fmt.Sprintf("%x", md5.Sum([]byte(in.InString)))}, nil
}
