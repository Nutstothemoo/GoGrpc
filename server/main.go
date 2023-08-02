package main

import (
	"log"
	"net"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)
type helloserver struct {
	pb.GreetServiceServer
}


func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	log.Printf("Starting gRPC server at %v ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", port, err)
	}

}