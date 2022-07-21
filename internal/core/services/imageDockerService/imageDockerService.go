package imageDockerService

import (
	"context"
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
	ctx := context.Background()
	rdb, _ := srv.redisRepository.GetRedisClient(ctx)

	if srv.redisRepository.ImageExist(rdb, image, tag) {
		return srv.redisRepository.Read(rdb, image, tag)
	}

	resp, _ := srv.dockerHubRepository.Read(rdb, image, tag)

	return resp, nil
}
