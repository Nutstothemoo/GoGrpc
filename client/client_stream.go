package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)
func callSayHelloClientStream (client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Client steaming RPC call with request: %v\n", names.Names)
	stream, err:= client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloClientStreaming RPC: %v", err)
	}
	for _, name := range names.Names {
		req:= &pb.HelloRequest{
			Name: name,
		}
		fmt.Printf("Sending req: %v\n", req)
		if err:=stream.Send(req); err != nil {
			log.Fatalf("Error while sending req: %v", err)
		}
		log.Printf("Response from SayHelloClientStreaming: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err :=stream.CloseAndRecv()
	log.Printf("Client steaming finished")
	if err != nil {
		log.Fatalf("Error while receiving response from SayHelloClientStreaming RPC: %v", err)
	}
	log.Printf("Response from SayHelloClientStreaming: %v", res.Messages)
}