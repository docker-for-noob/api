package imageReferenceService

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/docker-generator/api/internal/core/services/splitImageDockerService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type imageReferenceService struct {
	imageReferenceRepository ports.ImageReferenceRepository
	dockerHubRepository ports.DockerHubRepository
	imageDockerService ports.ImageDockerService
}

func New(imageReferenceRepository ports.ImageReferenceRepository, dockerHubRepository ports.DockerHubRepository, imageDockerService ports.ImageDockerService) *imageReferenceService {
	return &imageReferenceService{
		imageReferenceRepository: imageReferenceRepository,
		dockerHubRepository: dockerHubRepository,
		imageDockerService: imageDockerService,
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
			return errors.New(apperrors.Internal, err, "An internal error occured while searching ALL the reference", "")
		}
	}


	err := srv.imageReferenceRepository.AddAllTagReferenceFromApi(splitImageDockerService.SplitDockerImageName)
	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occured while adding the reference in dadabase", "")
	}
	return nil
}

func (srv *imageReferenceService) FindAllTagReferenceForALanguage(languageName string) error {

	result, err := srv.imageDockerService.Get(languageName, "")
	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occured while searching the tags", "")
	}
	err = srv.dockerHubRepository.HandleMultipleGetTagReference(languageName, result.Tags)
	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occured while searching the reference", "")
	}

	return nil
}

// non utilisé pour le moment, a utiliser lorsque l'on ne trouve pas d'info dans le mongo
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
