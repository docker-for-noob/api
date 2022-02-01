package ports

import (
	domain "github.com/docker-generator/api/internal/core/domain/dockerHubDomain"
	"net/http"
)

type DockerHubRepository interface {
	ReadAll() (*http.Response, error)
	Read(image string, tag string) (domain.DockerHubResult, error)
}

type DockerHubService interface {
	GetAll() (*http.Response, error)
	Get(image string, tag string) (domain.DockerHubResult, error)
}
