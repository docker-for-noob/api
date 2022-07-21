package imageDockerService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
)

type imageDockerService struct {
	dockerHubRepository ports.DockerHubRepository
	redisRepository     ports.RedisRepository
}

func New(dockerHubRepository ports.DockerHubRepository, redisRepository ports.RedisRepository) *imageDockerService {
	return &imageDockerService{
		dockerHubRepository: dockerHubRepository,
		redisRepository:     redisRepository,
	}
}

func (srv *imageDockerService) Get(image string, tag string) (domain.DockerImageResult, error) {

	if srv.redisRepository.ImageExist(image, tag) {
		return srv.redisRepository.Read(image, tag)
	}

	resp, _ := srv.dockerHubRepository.Read(image, tag)

	return resp, nil
}
