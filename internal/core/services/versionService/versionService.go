package versionService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type versionService struct{
	dockerComposeRepository ports.DockerComposeRepository
	versionRepository ports.VersionRepository
}

func New(dockerComposeRepository ports.DockerComposeRepository, versionRepository ports.VersionRepository) *versionService{
	return &versionService{
		dockerComposeRepository: dockerComposeRepository,
		versionRepository: versionRepository,
	}
}

func (srv *versionService) Add(dockerComposeId string) error  {

	previousVersion, readError := srv.dockerComposeRepository.Read(dockerComposeId)
	if readError != nil {
		if errors.Is(readError, apperrors.NotFound) {
			return errors.New(apperrors.NotFound, nil, "previous docker-compose version not found", "")
		}
		return errors.New(apperrors.Internal, nil, "An internal error occured while searching the pervious version", "")
	}

	versionError := srv.versionRepository.Create(previousVersion)
	if versionError != nil {
		return errors.New(apperrors.Internal, nil, "An internal error occured while creating the version", "")
	}

	return nil
}

func (srv *versionService) Get(dockerComposeId string, versionId string) (domain.DockerCompose, error)  {

	dockerComposeVersion, err := srv.versionRepository.Read(dockerComposeId, versionId)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "an error occured while searching the version", "")
	}

	if (domain.DockerCompose{}) == dockerComposeVersion {
		return domain.DockerCompose{}, errors.New(apperrors.NotFound, nil, "version not found", "")
	}

	return dockerComposeVersion, nil
}

func (srv *versionService) GetAll(dockerComposeId string) ([]domain.DockerCompose, error) {

	allVersions, err := srv.versionRepository.ReadAll(dockerComposeId)
	if err != nil {
		return []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "an error occured while searching the version", "")
	}

	if len(allVersions) == 0 {
		return []domain.DockerCompose{}, errors.New(apperrors.NotFound, nil, "version not found", "")
	}
	return allVersions, err
}