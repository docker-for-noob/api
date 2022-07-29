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

	resp, err := srv.dockerHubRepository.Read(image, tag)

	return resp, err
}

func (srv *imageDockerService) GetImages() (domain.DockerImagesParse, error) {

	resp, err := srv.dockerHubRepository.GetImages()

	return resp, err
}

func (srv *imageDockerService) GetAllVersionsFromImage(image string) (domain.DockerImageVersions, error) {

	resp, err := srv.dockerHubRepository.GetAllVersionsFromImage(image)

	return resp, err
}

func (srv *imageDockerService) GetAllTagsFromImageVersion(image string, version string) (domain.DockerImageDetails, error) {

	resp, err := srv.dockerHubRepository.GetAllTagsFromImageVersion(image, version)

	return resp, err
}
