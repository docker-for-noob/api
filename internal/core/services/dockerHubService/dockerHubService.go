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

func (srv *dockerHubService) Get(id string) (domain.DockerHub, error) {

	dockerHub, err := srv.dockerHubRepository.Read(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.DockerHub{}, errors.New(apperrors.NotFound, err, "Image not found in docker hub", "")
		}

		return domain.DockerHub{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")

	}

	return dockerHub, nil
}
