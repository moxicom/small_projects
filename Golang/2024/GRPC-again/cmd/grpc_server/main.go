package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-proj/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	pb.UnimplementedShippingServiceServer
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Got new request with id=%v\n", req.GetId())

	return &pb.GetResponse{
		Info: &pb.UserInfo{
			Id:      req.GetId(),
			Name:    "Dmitriy",
			IsPidor: true,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listed: %s", err.Error())
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	reflection.Register(s)
	pb.RegisterShippingServiceServer(s, &server{})

	log.Println("Serving gRPC on 0.0.0.0:" + fmt.Sprint(grpcPort))

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
