package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type DockerHubRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
	GetImages() (domain.DockerImages, error)
	GetTagReference(image string, tag string) (domain.ImageReference, error)
	HandleMultipleGetTagReference(languageName string, allTags []string) error
}

type RedisRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
	ImageExist(image string, tag string) bool
}

type ImageDockerService interface {
	Get(image string, tag string) (domain.DockerImageResult, error)
	GetImages() (domain.DockerImages, error)
}
