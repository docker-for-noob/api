package imageReferenceService

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"log"
)

type imageReferenceService struct {
	imageReferenceRepository ports.ImageReferenceRepository
	dockerHubRepository ports.DockerHubRepository
}

func New(imageReferenceRepository ports.ImageReferenceRepository, dockerHubRepository ports.DockerHubRepository) *imageReferenceService {
	return &imageReferenceService{
		imageReferenceRepository: imageReferenceRepository,
		dockerHubRepository: dockerHubRepository,
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

func (srv *imageReferenceService) AddAllTagReference(allLanguage []string) error {
	for _, languageName := range allLanguage {
		err := srv.FindAllTagReferenceForALanguage(languageName)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := srv.imageReferenceRepository.AddAllTagReferenceFromApi()
	if err != nil {
		return err
	}
	return nil
}

func (srv *imageReferenceService) FindAllTagReferenceForALanguage(languageName string) error {

	result, _ := srv.dockerHubRepository.Read(languageName, "")

	err := srv.dockerHubRepository.HandleMultipleGetTagReference(languageName, result.Tags)
	if err != nil {
		return err
	}

	return nil
}

func (srv *imageReferenceService) AddOneTagReferenceForALanguage(languageName string, tagName string) error {

	tagReference, err := srv.dockerHubRepository.GetTagReference(languageName , tagName )
	if err != nil {
		fmt.Printf("error on %s : %s\n",tagName, err)
	}
	err = srv.imageReferenceRepository.Add(tagReference)
	if err != nil {
		fmt.Printf("error on %s : %s\n",tagName, err)
	}

	return nil
}
