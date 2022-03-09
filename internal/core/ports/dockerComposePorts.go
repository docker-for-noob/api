package ports

import "github.com/docker-generator/api/internal/core/domain"

type DockerComposeRepository interface{
	ReadAll(firstItemRank int) ([]domain.DockerCompose, error )
	Read(id string) (domain.DockerCompose, error)
	Create(dockerCompose domain.DockerCompose, id string) (domain.DockerCompose, error)
	Update(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	Delete(id string) (bool, error)
}

type DockerComposeService interface{
	GetAll(firstItemRank int) (int, []domain.DockerCompose, error )
	Get(id string) (domain.DockerCompose, error)
	Post(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	Patch(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	Delete(id string) (bool, error)
}