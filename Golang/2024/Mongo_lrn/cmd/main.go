package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mongo_lrn/internal/config"
	"mongo_lrn/internal/models"
	"mongo_lrn/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	// log := initLogger()
	log.Println("Starting server")

	cfg := config.Get()
	log.Printf("Config got successfully. URI = %v", cfg.Uri)

	log.Println("Connecting to mongodb")
	s, ctx, cancelCtx := storage.New(cfg)
	log.Println("Connected to mongodb")

	defer cancelCtx()
	defer s.Client.Disconnect(ctx)

	defer func() {
		if err := s.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// insertedObj, err := s.Colls.Categories.InsertOne(ctx, models.Category{
	// 	Name: "Огнетуши",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("Inserted ID = %v", insertedObj.InsertedID)

	id, err := primitive.ObjectIDFromHex("6649c3f11aaa1ddbd696895b")
	if err != nil {
		panic(err)
	}

}

func find(ctx context.Context, s *storage.Storage) {
	cursor, err := s.Colls.Categories.Find(ctx, bson.D{{"name", "Огнетуши"}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			panic(fmt.Errorf("no document found err $(%v)", err))
		} else {
			panic(err)
		}
	}

	var categories []models.Category
	if err := cursor.All(ctx, &categories); err != nil {
		panic(err)
	}

	for _, category := range categories {
		log.Println(category)
	}
}
