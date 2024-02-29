package services

import (
	"context"
	"log"
	"redis/models"
	"redis/repository/postgres"
	"redis/repository/redisRep"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	r redisRep.RedisClient
	p postgres.Postgres
}

func NewService(redisCl redisRep.RedisClient, db postgres.Postgres) *Service {
	return &Service{redisCl, db}
}

func (s *Service) GetUser(key string) (models.User, error) {
	// Redis
	ctx := context.Background()
	userRedis, err := s.r.GetUser(ctx, key)
	if err != nil && err != redis.Nil {
		return models.User{}, err
	}

	if err == nil && userRedis != (models.User{}) {
		log.Printf("User found in Redis: %v", userRedis)
		return userRedis, nil
	}

	// Postgres
	userPostgres, err := s.p.GetUser(key)
	if err != nil {
		return models.User{}, err
	}
	log.Printf("User found in Postgres")

	// Set the user in Redis
	err = s.r.SetUser(ctx, key, userPostgres)
	if err != nil {
		log.Printf("Error setting user in Redis: %v", err)
		// Don't return error here since user is successfully retrieved from Postgres
	}

	return userPostgres, nil
}
