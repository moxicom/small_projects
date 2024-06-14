package storage

import (
	"context"
	"elastic/internal/pkg/models"
)

type ProductStorer interface {
	Insert(ctx context.Context, prod models.Product) error
	Update(ctx context.Context, prod models.Product) error
	Delete(ctx context.Context, prod models.Product) error
	FindMany(ctx context.Context, prod models.Product) ([]models.Product, error)
}
