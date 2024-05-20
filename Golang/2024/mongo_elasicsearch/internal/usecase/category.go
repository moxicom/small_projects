package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_lrn/internal/models"
)

func (uc *Usecase) NewCategory(name string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.TODO(), contexTimeOut)
	newCategory := models.Category{Name: name}

	insertedID, err := uc.Repo.Category.NewCategory(ctx, newCategory)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return insertedID, err
}

func (uc *Usecase) GetCategories() ([]models.Category, error) {
	ctx, _ := context.WithTimeout(context.TODO(), contexTimeOut)
	return uc.Repo.Category.GetCategories(ctx)
}

func (uc *Usecase) GetCategoryByID(id interface{}) (models.Category, error) {
	ctx, _ := context.WithTimeout(context.TODO(), contexTimeOut)
	return uc.Repo.Category.GetCategoryByID(ctx, id)
}

func (uc *Usecase) DeleteCategories() error {
	ctx, _ := context.WithTimeout(context.TODO(), contexTimeOut)
	return uc.Repo.Category.DeleteCategories(ctx)
}

func (uc *Usecase) DeleteCategoryByID(id string) error {
	ctx, _ := context.WithTimeout(context.TODO(), contexTimeOut)
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	return uc.Repo.Category.DeleteCategoryByID(ctx, oid)
}
