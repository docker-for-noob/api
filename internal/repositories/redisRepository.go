package repositories

import (
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-redis/redis/v8"
)

type redisRepository struct{}

func NewRedisRepository() *redisRepository {
	return &redisRepository{}
}

func GetClient(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     goDotEnv.GetEnvVariable("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}

func (repo *redisRepository) Read(image string, tag string) (domain.DockerImageResult, error) {
	ctx := context.Background()
	rdb, _ := repo.GetRedisClient(ctx)

	var dockerHubTags []string

	dockerHubTags = rdb.LRange(ctx, image+"-"+tag, 0, -1).Val()

	DockerImageResult := domain.DockerImageResult{
		Name: image,
		Tags: dockerHubTags,
	}

	return DockerImageResult, nil

}

func (repo *redisRepository) AddImage(image string, tag string, encoded []byte) {
	ctx := context.Background()
	rdb, _ := repo.GetRedisClient(ctx)
	rdb.RPush(ctx, image+"-"+tag, encoded)
}

func (repo *redisRepository) ImageExist(image string, tag string) bool {
	ctx := context.Background()

	rdb, err := repo.GetRedisClient(ctx)

	if err != nil {
		panic(err)
	}

	length := rdb.LLen(ctx, image+"-"+tag).Val()

	return length > 0
}
