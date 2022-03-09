package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func GetClient(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Failed to connect redis: %v", err)
	}

	return rdb
}
