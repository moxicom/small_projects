package usecase

import (
	"mongo_lrn/internal/storage/mongo_storage"
)

type Usecase struct {
	Repo *mongo_storage.Repository
}

func New(r *mongo_storage.Repository) *Usecase {
	return &Usecase{r}
}
