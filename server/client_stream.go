package main

import (
	"io"
	"log"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func (s *helloserver)SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the client stream.
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Received request with name: %v", req.Name)
		messages = append(messages,"Hello", req.Name)
	}
}