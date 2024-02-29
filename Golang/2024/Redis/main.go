package main

import (
	"context"
	"fmt"
	"log"
	"redis/repository/postgres"
	"redis/repository/redisRep"
	"redis/services"
	"time"

	"github.com/lpernett/godotenv"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s", err)
	}

	db, err := postgres.NewDbInit()
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&models.User{})

	repo := postgres.NewRepository(db)

	ctx := context.Background()

	redisCl := redisRep.NewRedisInit()
	defer redisCl.Client.Close()

	if err := redisCl.Client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	if err = redisCl.Client.Set(ctx, "asd", "qwe", time.Second*10).Err(); err != nil {
		panic(err)
	}

	service := services.NewService(redisCl, repo)

	user, err := service.GetUser("user:12")
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}

	log.Println(user)

	fmt.Println(repo, redisCl)
}
