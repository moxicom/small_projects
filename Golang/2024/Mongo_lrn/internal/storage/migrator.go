package storage

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

type Collections struct {
	Categories *mongo.Collection
	Products   *mongo.Collection
}

func (s *Storage) migrateCollections(ctx context.Context) error {
	log.Println("Migration started")

	dbName := "testdb"

	collections := []string{CategoriesCol, ProductsCol}

	for _, collName := range collections {
		coll := s.Client.Database(dbName).Collection(collName)
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

	databases, err := s.Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return err
	}

	log.Println(databases)

	// Add collections to struct
	s.Colls.Categories = s.Client.Database(dbName).Collection(CategoriesCol)
	s.Colls.Categories = s.Client.Database(dbName).Collection(ProductsCol)

	return nil
}
