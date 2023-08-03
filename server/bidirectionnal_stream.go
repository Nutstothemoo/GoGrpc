package main

import (
	"io"
	"log"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func (s *helloserver) SayHelloBidirectionalStreamingServer( stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err:= stream.Recv()
		if err != io.EOF {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		if err!= nil {	
			log.Fatalf("Error while reading client stream: %v", err)
		}
		log.Printf("Got request with : %v", req.Name)
		res:= &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err:= stream.Send(res); err != nil {
			log.Fatalf("Error while sending response: %v", err)
		}
	}
}

