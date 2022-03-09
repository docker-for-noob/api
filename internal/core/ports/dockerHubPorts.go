package ports

import "github.com/docker-generator/api/internal/core/domain"

type DockerHubRepository interface {
	Read(image string, tag string) (domain.DockerHubResult, error)
}

type DockerHubService interface {
	Get(image string, tag string) (domain.DockerHubResult, error)
}
