package main

import (
	"context"
	"io"
	"log"

	pb "github.com/taufiqoo/learn-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
