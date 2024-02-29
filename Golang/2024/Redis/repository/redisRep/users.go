package redisRep

import (
	"context"
	"log"
	"redis/models"
	"time"

	"github.com/redis/go-redis/v9"
)

func (c *RedisClient) GetUser(ctx context.Context, key string) (models.User, error) {
	var user models.User
	err := c.Client.HGetAll(ctx, key).Scan(&user)
	return user, err
}

func (c *RedisClient) SetUser(ctx context.Context, key string, user models.User) error {
	log.Printf("inserting to redis: %v\n", user)
	_, err := c.Client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "id", user.Id)
		rdb.HSet(ctx, key, "name", user.Name)
		rdb.HSet(ctx, key, "age", user.Age)
		return nil
	})

	if err != nil {
		return err
	}

	err = c.Client.Expire(ctx, key, time.Second*2).Err()
	return err
}
