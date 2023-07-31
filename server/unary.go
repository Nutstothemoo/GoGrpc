package main

import (
	"context"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)

func(s *helloserver) SayHello( ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " }, nil
}