package service

import (
	"context"
	"elastic/internal/pkg/models"
	"elastic/internal/pkg/storage"
	"encoding/json"
	"fmt"
	"strconv"
)

type Service struct {
	Storage storage.ProductStorer
	Pg      interface{}
}

func New(storage storage.ProductStorer) *Service {
	return &Service{Storage: storage}
}

func (s *Service) InsertProduct(ctx context.Context, prod models.Product) error {
	id := strconv.Itoa(prod.ID)
	jsonBody, err := json.Marshal(prod)
	if err != nil {
		return fmt.Errorf("error on marshalling json body")
	}

	return s.Storage.InsertProduct(ctx, id, jsonBody)
}

func (s *Service) UpdateProduct(ctx context.Context, prod models.Product) error {
	id := strconv.Itoa(prod.ID)
	return s.Storage.UpdateProduct(ctx, id, prod)
}

// func (s *Service) DeleteProduct(ctx context.Context, prodID int) error {
// 	return s.Storage.DeleteProduct(ctx, prodID)
// }

func (s *Service) SearchProduct(ctx context.Context, searchStr string) ([]models.Product, error) {
	return s.Storage.FindManyProducts(ctx, searchStr)
}
