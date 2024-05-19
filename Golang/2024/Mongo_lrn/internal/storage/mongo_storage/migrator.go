package mongo_storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CategoriesCol = "categories"
	ProductsCol   = "products"
)

func migrateCollections(ctx context.Context, db *mongo.Database) error {
	log.Println("Migration started")

	collections := []string{CategoriesCol, ProductsCol}

	for _, collName := range collections {
		coll := db.Collection(collName)
		count, err := coll.EstimatedDocumentCount(ctx)
		if err != nil {
			return err
		}

		if count == 0 {
			_, err = coll.InsertOne(ctx, bson.M{"empty_init_value": "empty"})
			if err != nil {
				return err
			}
			log.Printf("Inserted dummy document into collection %s", collName)
		} else {
			log.Printf("Collection %s is already exists", collName)
		}
	}

	return nil
}
