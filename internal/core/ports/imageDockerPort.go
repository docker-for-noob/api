package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type DockerHubRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
	GetImages() (domain.DockerImagesParse, error)
	GetAllTagsFromImageVersion(image string, version string) (domain.ImageNameDetail, error)
	GetTagReference(image string, tag string) (domain.ImageReference, error)
	HandleMultipleGetTagReference(languageName string, allTags []string) error
}

type RedisRepository interface {
	Read(image string, tag string) (domain.DockerImageResult, error)
	ImageExist(image string, tag string) bool
	Add(key string, value interface{})
	FindDockerImageResult(key string) []string
}

type ImageDockerService interface {
	Get(image string, tag string) (domain.DockerImageResult, error)
	GetImages() (domain.DockerImagesParse, error)
	GetAllVersionsFromImage(image string) (domain.DockerImageVersions, error)
	GetAllTagsFromImageVersion(image string, version string) ([]domain.ImageNameDetail, error)
}
