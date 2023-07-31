package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Nutstothemoo/GoGrpc/proto"
)
func callSayHello(client pb.GreetServiceClient) {
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

res, err := client.SayHello(ctx, &pb.NoParams{})
if err != nil {
	log.Fatalf("Error when calling SayHello: %v", err)
}
log.Printf("Response from server: %v", res.Message)
}