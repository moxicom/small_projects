package category_rep

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CategoryRep) DeleteCategories(ctx context.Context) error {
	filter := bson.D{}

	_, err := r.col.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRep) DeleteCategoryByID(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}

	fmt.Println(filter)
	_, err := r.col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
