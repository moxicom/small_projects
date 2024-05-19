package storage

import (
	"context"
	"mongo_lrn/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Storage struct {
	Client *mongo.Client
	Colls  Collections
}

func New(cfg config.Config) (Storage, context.Context, context.CancelFunc) {
	ctx, client, cancel := MustSetupMongoDB(cfg)

	storage := Storage{client, Collections{}}

	if err := storage.migrateCollections(ctx); err != nil {
		panic(err)
	}

	return storage, ctx, cancel
}

func MustSetupMongoDB(cfg config.Config) (context.Context, *mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return ctx, client, cancel
}
