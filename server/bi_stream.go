package main

import (
	"io"
	"log"

	pb "github.com/taufiqoo/learn-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request name with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
