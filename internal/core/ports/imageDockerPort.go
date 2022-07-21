package ports

import (
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/go-redis/redis/v8"
)

type DockerHubRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
}

type RedisRepository interface {
	GetRedisClient(ctx context.Context) (*redis.Client, error)
	Read(image string, tag string) (domain.DockerImageResult, error)
	//AddImage(image string, tag string) (domain.DockerImageResult, error)
	ImageExist(image string, tag string) bool
}

type ImageDockerService interface {
	Get(image string, tag string) (domain.DockerImageResult, error)
}
