package category_rep

import (
	"context"
	"errors"
	"mongo_lrn/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CategoryRep) NewCategory(ctx context.Context, newCategory models.Category) (oid primitive.ObjectID, err error) {
	result, err := r.col.InsertOne(ctx, newCategory)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("invalid oid")
	}

	return oid, nil
}
