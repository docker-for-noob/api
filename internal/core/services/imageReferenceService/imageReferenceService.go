package imageReferenceService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
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
	return domain.ImageReference{}, nil
}