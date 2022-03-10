package repositories

import (
	"context"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-redis/redis/v8"
)

func GetClient(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     goDotEnv.GetEnvVariable("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}
