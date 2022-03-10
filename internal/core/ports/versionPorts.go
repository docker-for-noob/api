package ports

import "github.com/docker-generator/api/internal/core/domain"

type VersionRepository interface {
	Create(dockerCompose domain.DockerCompose, userId string) error
	Read(id string, idVersion string, userId string) (domain.DockerCompose, error)
	ReadAll(id string, userId string) ([]domain.DockerCompose, error)
}

type VersionService interface {
	Add(dockerComposeId string, userId string) error
	Get(id string, idVersion string, userId string) (domain.DockerCompose, error)
	GetAll(id string, userId string) ([]domain.DockerCompose, error)
}
