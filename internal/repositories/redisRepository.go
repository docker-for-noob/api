package repositories

import (
	"context"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisRepository() *redisRepository {
	ctx := context.Background()
	rdb, _ := GetRedisClient(ctx)
	return &redisRepository{rdb: rdb, ctx: ctx}
}

func GetRedisClient(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     goDotEnv.GetEnvVariable("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}

func (repo *redisRepository) Read(image string, tag string) (domain.DockerImageResult, error) {
	var dockerHubTags []string

	dockerHubTags = repo.rdb.LRange(repo.ctx, image+"-"+tag, 0, -1).Val()

	DockerImageResult := domain.DockerImageResult{
		Name: image,
		Tags: dockerHubTags,
	}

	return DockerImageResult, nil

}

func (repo *redisRepository) ImageExist(image string, tag string) bool {
	length := repo.rdb.LLen(repo.ctx, image+"-"+tag).Val()

	return length > 0
}

func (repo *redisRepository) Add(key string, value interface{}) {
	listToString, _ := json.Marshal(value)
	if repo.rdb.LLen(repo.ctx, key).Val() > 0 {
		repo.rdb.Del(repo.ctx, key)
	}
	
	repo.rdb.RPush(repo.ctx, key, listToString)
}

func (repo *redisRepository) FindDockerImageResult(key string) []string {
	if repo.rdb.LLen(repo.ctx, key).Val() > 0 {
		return repo.rdb.LRange(repo.ctx, key, 0, -1).Val()
	}
	return []string{}
}


