package ports

import "github.com/docker-generator/api/internal/core/domain"

type DockerComposeRepository interface{
	Read(id string) (domain.DockerCompose, error)
	Create(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	//Update(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	//Delete(id string) (bool, error)
}

type DockerComposeService interface{
	Get(id string) (domain.DockerCompose, error)
	Post(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	//Patch(dockerCompose domain.DockerCompose) (domain.DockerCompose, error)
	//Delete(id string) (bool, error)
}

