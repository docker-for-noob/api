package ports

import "github.com/docker-generator/api/internal/core/domain"

type VersionRepository interface{
	Create(dockerCompose domain.DockerCompose) error
	Read(id string, idVersion string) (domain.DockerCompose, error)
	ReadAll(id string) ([]domain.DockerCompose, error)
}


type VersionService interface {
	Add(dockerComposeId string) error
	Get(id string, idVersion string) (domain.DockerCompose, error)
	GetAll(id string) ([]domain.DockerCompose, error)
}