package repositories

import "github.com/docker-generator/api/internal/core/domain"

type imageReferenceRepository struct {}

func NewImageReferenceRepository() *imageReferenceRepository {
	return &imageReferenceRepository{}
}

func (repository *imageReferenceRepository) Read(imageName string) (domain.ImageReference, error) {
	return domain.ImageReference{}, nil
}