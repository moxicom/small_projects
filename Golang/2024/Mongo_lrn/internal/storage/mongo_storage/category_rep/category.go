package category_rep

import "go.mongodb.org/mongo-driver/mongo"

type CategoryRep struct {
	col *mongo.Collection
}

func NewCategoryRep(col *mongo.Collection) *CategoryRep {
	return &CategoryRep{col}
}
