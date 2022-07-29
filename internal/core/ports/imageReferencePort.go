package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/repositories"
)

type ImageReferenceRepository interface {
	Read(imageName string) (domain.ImageReference, error)
	Add(imageReference domain.ImageReference) error
	AddAllTagReferenceFromApi(fn repositories.Formater) error
}

type ImageReferenceService interface {
	Get(imageName string) (domain.ImageReference, error)
	FindAllTagReferenceForALanguage(languageName string) error
	AddAllTagReference(allLanguage []string) error
}
