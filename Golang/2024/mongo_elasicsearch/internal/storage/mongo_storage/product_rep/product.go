package product_rep

import "go.mongodb.org/mongo-driver/mongo"

type ProductRep struct {
	col *mongo.Collection
}

func NewProductRep(col *mongo.Collection) *ProductRep {
	return &ProductRep{col}
}
