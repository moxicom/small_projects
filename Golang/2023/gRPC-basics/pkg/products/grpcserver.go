package products

import (
	"context"
	grpc_basics "gprc_basics/pkg/api/grpc-basics"
)

// GRPCServer ...
type GRPCServer struct{}

// Return value ...
func (s *GRPCServer) AddProduct(context.Context, *grpc_basics.Product) (*grpc_basics.ProductID, error) {
	return &grpc_basics.ProductID{Value: "z4df5gkkdpalj53k5"}, nil
}
