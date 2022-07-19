package imageReferenceService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type imageReferenceService struct {
	imageReferenceRepository ports.ImageReferenceRepository
}

func New(imageReferenceRepository ports.ImageReferenceRepository) *imageReferenceService {
	return &imageReferenceService{
		imageReferenceRepository: imageReferenceRepository,
	}
}

func (srv *imageReferenceService) Get(imageName string) (domain.ImageReference, error) {

	result, err := srv.imageReferenceRepository.Read(imageName)

	if err != nil {
		return domain.ImageReference{}, errors.New(apperrors.Internal, err, "An internal error occured while searching the reference", "")
	}

	if "" == result.Name  {
		return result, errors.New(
			apperrors.NotFound,
			nil,
			"Not found",
			"",
		)
	}

	return result, nil
}