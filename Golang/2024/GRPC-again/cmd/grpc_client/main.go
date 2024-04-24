package main

import (
	"context"
	"log"
	"time"

	pb "grpc-proj/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "localhost:50051"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Dial error: %s", err.Error())
	}
	defer conn.Close()

	// Create a client
	c := pb.NewShippingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make a request with name 'Get'
	r, err := c.Get(ctx, &pb.GetRequest{Id: 1})
	if err != nil {
		log.Fatalf("Request error: %s", err.Error())
	}

	log.Printf("Response: %v", r.GetInfo())
}
