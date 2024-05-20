package main

import (
	"context"
	"log"
	"mongo_lrn/internal/config"
	"mongo_lrn/internal/storage/mongo_storage"
	"mongo_lrn/internal/usecase"
)

func main() {

	// log := initLogger()
	log.Println("Starting server")

	cfg := config.Get()
	log.Printf("Config got successfully. URI = %v", cfg.Uri)

	log.Println("Connecting to mongodb")
	client, ctx, cancelCtx := mongo_storage.Connect(cfg)
	log.Println("Connected to mongodb")

	defer cancelCtx()
	defer client.Disconnect(ctx)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := mongo_storage.NewRepository(client)
	uc := usecase.New(r)

	//id, err := uc.NewCategory("zxcqwe")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(id)

	categories, err := uc.GetCategories()
	if err != nil {
		panic(err)
	}
	log.Println(categories)

	//category, err := uc.GetCategoryByID(id)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(category)

	err = uc.DeleteCategoryByID("664b1e14c35388c7fdc4acfe")
	if err != nil {
		panic(err)
	}
}
