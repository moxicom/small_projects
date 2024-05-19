package category_rep

import (
	"context"
	"mongo_lrn/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CategoryRep) NewCategory(ctx context.Context, newCategory models.Category) (insertedID interface{}, err error) {
	result, err := r.col.InsertOne(ctx, newCategory)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	insertedID = result.InsertedID

	return insertedID, nil
}
