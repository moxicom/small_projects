package storage

import (
	"context"
	"elastic/internal/pkg/models"
	"elastic/internal/pkg/storage/elastic"
)

type ProductStorer interface {
	InitIndexes(indexes elastic.Indexes) error
	InsertProduct(ctx context.Context, prod models.Product) error
	UpdateProduct(ctx context.Context, prod models.Product) error
	DeleteProduct(ctx context.Context, prodID int) error
	FindManyProducts(ctx context.Context, searchStr string) ([]models.Product, error)

	//InsertCategory(ctx context.Context, prod models.Product) error
	//UpdateCategory(ctx context.Context, prod models.Product) error
	//DeleteCategory(ctx context.Context, prod models.Product) error
	//FindManyCategories(ctx context.Context, prod models.Product) ([]models.Product, error)
}
