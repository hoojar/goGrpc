package api

import (
	"context"
	"fmt"

	pb "grpc/myGrpc"
)

func (s *rpcService) MathSum(ctx context.Context, in *pb.NumRequest) (*pb.NumResponse, error) {
	fmt.Printf("Math Sum Min: %d max: %d \n", in.Min, in.Max)
	return &pb.NumResponse{Amount: in.Min + in.Max}, nil
}
