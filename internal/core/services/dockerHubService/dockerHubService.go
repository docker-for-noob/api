package dockerHubService

import (
	domain "github.com/docker-generator/api/internal/core/domain/dockerHubDomain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"log"
	"net/http"
)

type dockerHubService struct {
	dockerHubRepository ports.DockerHubRepository
}

func New(dockerHubRepository ports.DockerHubRepository) *dockerHubService {
	return &dockerHubService{
		dockerHubRepository: dockerHubRepository,
	}
}

func (srv *dockerHubService) GetAll() (*http.Response, error) {

	resp, err := srv.dockerHubRepository.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return resp, nil
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
