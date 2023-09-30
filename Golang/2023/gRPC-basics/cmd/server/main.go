package main

import (
	grpc_basics "gprc_basics/pkg/api/grpc-basics"
	"gprc_basics/pkg/products"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &products.GRPCServer{}
	grpc_basics.RegisterProductInfoServer(s, srv)

	// l, err := net.Listen("tcp", ":8080")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := s.Serve(l); err != nil {
	// 	log.Fatal(err)
	// }

}
