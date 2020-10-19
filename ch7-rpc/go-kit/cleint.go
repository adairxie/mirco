package main

import (
	"context"
	"flag"
	"fmt"
	service "rpc/go-kit/string-service"
	"rpc/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()

	svr := NewStringClient(conn)
	result, err := svr.Concat(ctx, "A", "B")
	if err != nil {
		fmt.Println("Check error", err.Error())
	}

	fmt.Println("result =", result)
}

func NewStringClient(conn *grpc.ClientConn) service.Service {
	var ep = grpctransport.NewClient(conn,
		"pb.StringService",
		"Concat",
		EncodeStringRequest,
		DecodeStringResponse,
		pb.StringResponse{},
	).Endpoint()

	userEp := service.StringEndpoints{
		StringEndpoint: ep,
	}

	return userEp
}

func DecodeStringResponse(ctx context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func EncodeStringRequest(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}
