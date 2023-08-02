package main

import (
	"log"
	"time"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func (s *helloserver) SayHelloServerStreaming( req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received request from client: %v", req.Names)
	for _, name := range req.Names {
		res:= &pb.HelloResponse{
			Message: "Hello " + name		}
	
		if err:= stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}