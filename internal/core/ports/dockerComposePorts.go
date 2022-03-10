package ports

import "github.com/docker-generator/api/internal/core/domain"

type DockerComposeRepository interface {
	ReadAll(firstItemRank int, userId string) ([]domain.DockerCompose, error)
	Read(id string, userId string) (domain.DockerCompose, error)
	Create(dockerCompose domain.DockerCompose, userId string, id string) (domain.DockerCompose, error)
	Update(dockerCompose domain.DockerCompose, userId string) (domain.DockerCompose, error)
	Delete(id string, userId string) (bool, error)
}

type DockerComposeService interface {
	GetAll(firstItemRank int, userId string) (int, []domain.DockerCompose, error)
	Get(id string, userId string) (domain.DockerCompose, error)
	Post(dockerCompose domain.DockerCompose, userId string) (domain.DockerCompose, error)
	Patch(dockerCompose domain.DockerCompose, userId string) (domain.DockerCompose, error)
	Delete(id string, userId string) (bool, error)

}