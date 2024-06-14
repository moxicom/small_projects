package service

import "elastic/internal/pkg/storage"

type Service struct {
	storage storage.ProductStorer
}

func New(storage storage.ProductStorer) *Service {
	return &Service{storage}
}

func (s *Service)
