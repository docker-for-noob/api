package imageReferenceService

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
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

func (srv *imageReferenceService) AddAllTagReferenceForALanguage(languageName string) error {
	//var wg sync.WaitGroup

	result, _ := srv.dockerHubRepository.Read(languageName, "") //175 * 1000 ==> 175000

	err := srv.dockerHubRepository.HandleMultipleGetTagReference(languageName, result.Tags)
	if err != nil {
		return err
	}

	// lenght total des elements
	// var requestcount
	//var requestTimer
	// boucle sur tant que LEnght - count != 0

		// repo.FetchMultipleAsync(Count, {})
		// wait timer + 5s

	/*for _ , tag := range result.Tags{
		err := srv.AddOneTagReferenceForALanguage(languageName, tag)
		if err != nil {
			fmt.Printf("%s - %s\n", tag, err)
		}
	}*/

	/*for _ , tag := range result.Tags{
		wg.Add(1)
		go func(tag string) {
			defer wg.Done()
			err := srv.AddOneTagReferenceForALanguage(languageName, tag)
			if err != nil {
				fmt.Printf("%s - %s\n", tag, err)
			}
		}(tag)
	}
	wg.Wait()*/
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

	//fmt.Printf("%s - %s\n",tagName, tagReference)

	return nil
}
