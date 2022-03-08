package dockerComposeService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type dockerComposeService struct {
	dockerComposeRepository ports.DockerComposeRepository
	versionService ports.VersionService
}

func New(dockerComposeRepository ports.DockerComposeRepository, versionService ports.VersionService) *dockerComposeService {
	return &dockerComposeService{
		dockerComposeRepository: dockerComposeRepository,
		versionService: versionService,
	}
}

func (srv *dockerComposeService) GetAll(firstItemRank int) (int, []domain.DockerCompose, error ) {

	dockerComposeList, err := srv.dockerComposeRepository.ReadAll(firstItemRank)
	if err != nil {
		return 0, []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	lastItemRank :=  firstItemRank + 25
	return lastItemRank, dockerComposeList, nil
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
func (srv *dockerComposeService) Post(dockerCompose domain.DockerCompose) (domain.DockerCompose, error){

	dockerComposeResult, err := srv.dockerComposeRepository.Create(dockerCompose)

	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return dockerComposeResult, nil
}

func (srv *dockerComposeService) Patch(dockerCompose domain.DockerCompose) (domain.DockerCompose, error){

	versionErr := srv.versionService.Add(dockerCompose.Id)

	if versionErr != nil {
		if errors.Is(versionErr, apperrors.NotFound) {
			return domain.DockerCompose{}, errors.New(apperrors.NotFound, versionErr, "version Service can not found dockerCompose", "")
		}
		return domain.DockerCompose{}, errors.New(apperrors.Internal, versionErr, "An internal error occurred in versionService", "")
	}

	dockerComposeResult, err := srv.dockerComposeRepository.Update(dockerCompose)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.DockerCompose{}, errors.New(apperrors.NotFound, err, "DockerCompose not found in database", "")
		}

		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")

	}

	return dockerComposeResult, nil
}

func (srv *dockerComposeService) Delete(id string) (bool, error){

	isDeleted, err := srv.dockerComposeRepository.Delete(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return isDeleted, errors.New(apperrors.NotFound, err, "DockerCompose not found in database", "")
		}

		return isDeleted, errors.New(apperrors.Internal, err, "An internal error occurred", "")

	}

	return isDeleted, nil
}