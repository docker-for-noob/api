package ports

import "github.com/docker-generator/api/internal/core/domain"

type DockerHubRepository interface {
	Read(id string) (domain.DockerHub, error)
}

type DockerHubService interface {
	Get(id string) (domain.DockerHub, error)
}
