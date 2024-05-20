package category_rep

import (
	"context"
	"mongo_lrn/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *CategoryRep) GetCategories(ctx context.Context) ([]models.Category, error) {
	filter := bson.D{}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func (r *CategoryRep) GetCategoryByID(ctx context.Context, id interface{}) (models.Category, error) {
	filter := bson.D{{"_id", id}}

	var category models.Category
	err := r.col.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}
