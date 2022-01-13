package dockerComposeService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type dockerComposeService struct {
	dockerComposeRepository ports.DockerComposeRepository
}

func New(dockerComposeRepository ports.DockerComposeRepository) *dockerComposeService {
	return &dockerComposeService{
		dockerComposeRepository: dockerComposeRepository,
	}
}

func (srv *dockerComposeService) Get(id string) (domain.DockerCompose, error) {

	dockerCompose, err := srv.dockerComposeRepository.Read(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.DockerCompose{}, errors.New(apperrors.NotFound, err, "DockerCompose not found in database", "")
		}

		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")

	}

	return dockerCompose, nil
}