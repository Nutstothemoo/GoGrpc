package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func callHelloBidirectionalStreamingClient(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Bidirectional streaming RPC START")
	stream, err:= client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloBidirectionalStreaming RPC: %v", err)
	}
	waitchannel := make(chan struct{})
	go func(){
		for {
			message, err:= stream.Recv()
			if err == io.EOF {
				// fin du str
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)			
			}
			log.Println("Received: ", message.Message)
		}
	}()
	for _, name := range names.Names {
		req:= &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitchannel
	log.Printf("Bidirectional streaming RPC END")
}