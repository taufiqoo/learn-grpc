package main

import (
	"context"

	pb "github.com/taufiqoo/learn-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
