package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type DockerHubRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
}

type RedisRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
	//AddImage(image string, tag string) (domain.DockerImageResult, error)
	ImageExist(image string, tag string) bool
}

type ImageDockerService interface {
	Get(image string, tag string) (domain.DockerImageResult, error)
}
