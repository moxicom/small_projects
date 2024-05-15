package main

import (
	"context"
	"log"
	"mongo_lrn/internal/config"
	"mongo_lrn/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	// log := initLogger()
	log.Println("Starting server")

	cfg := config.Get()
	log.Printf("Config got successfully. URI = %v", cfg.Uri)

	log.Println("Connecting to mongodb")
	_, client, cancelCtx := storage.MustSetupMongoDB(cfg)
	log.Println("Connected to mongodb")

	defer cancelCtx()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

// func initLogger() slog.Logger {
// 	var log slog.Logger
// 	return log
// }
