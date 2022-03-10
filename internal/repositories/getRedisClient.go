package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"github.com/docker-generator/api/pkg/goDotEnv"
)

func GetClient(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     goDotEnv.GetEnvVariable("REDIS_URL"),
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Failed to connect redis: %v", err)
	}

	return rdb
}
