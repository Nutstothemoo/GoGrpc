package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names  *pb.NamesList) {
	log.Printf("Calling SayHelloServerStreaming RPC...")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Error when calling SayHelloServerStreaming: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream.
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from SayHelloServerStreaming: %v", message)
	}
	log.Printf("Finished SayHelloServerStreaming RPC...")
}