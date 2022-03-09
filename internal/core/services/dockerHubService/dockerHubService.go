package dockerHubService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type dockerHubService struct {
	dockerHubRepository ports.DockerHubRepository
}

func New(dockerHubRepository ports.DockerHubRepository) *dockerHubService {
	return &dockerHubService{
		dockerHubRepository: dockerHubRepository,
	}
}

func (srv *dockerHubService) Get(image string, tag string) (domain.DockerHubResult, error) {
	resp, err := srv.dockerHubRepository.Read(image, tag)

	if err != nil {
		return resp, errors.New(
			apperrors.NotFound,
			nil,
			"Not found",
			"",
		)
	}

	return resp, nil
}
