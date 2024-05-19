package usecase

import (
	"context"
	"mongo_lrn/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *Usecase) NewCategory(name string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*10)
	newCategory := models.Category{Name: name}

	insertedID, err := uc.Repo.Category.NewCategory(ctx, newCategory)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return insertedID, err
}
