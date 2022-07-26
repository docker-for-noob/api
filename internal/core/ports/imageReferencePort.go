package ports

import "github.com/docker-generator/api/internal/core/domain"

type ImageReferenceRepository interface {
	Read(imageName string) (domain.ImageReference, error)
	Add(imageReference domain.ImageReference) error
}

type ImageReferenceService interface {
	Get(imageName string) (domain.ImageReference, error)
	AddAllTagReferenceForALanguage(languageName string) error
}
