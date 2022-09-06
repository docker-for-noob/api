package imageDockerService

import (
	"encoding/json"
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/docker-generator/api/internal/core/services/splitImageDockerService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/docker-generator/api/pkg/sliceUtils"
	"github.com/matiasvarela/errors"
)

type imageDockerService struct {
	dockerHubRepository ports.DockerHubRepository
	redisRepository     ports.RedisRepository
	dbRepository		ports.ImageReferenceRepository
}

func New(dockerHubRepository ports.DockerHubRepository, redisRepository ports.RedisRepository, dbRepository	ports.ImageReferenceRepository) *imageDockerService {
	return &imageDockerService{
		dockerHubRepository: dockerHubRepository,
		redisRepository:     redisRepository,
		dbRepository: dbRepository,
	}
}

func (srv *imageDockerService) Get(image string, tag string) (domain.DockerImageResult, error) {
	if srv.redisRepository.ImageExist(image, tag) {
		return srv.redisRepository.Read(image, tag)
	}

	resp, err := srv.dockerHubRepository.Read(image, tag)

	return resp, err
}

func (srv *imageDockerService) GetImages() (domain.DockerImagesParse, error) {

	resp, err := srv.dockerHubRepository.GetImages()

	return resp, err
}

func (srv *imageDockerService) GetAllVersionsFromImage(languageName string) (domain.DockerImageVersions, error) {

	allImageForOnelanguage, err := srv.Get(languageName, "")
	if err != nil {
		return domain.DockerImageVersions{}, errors.New(apperrors.Internal, err, "An internal error occured while searching the tags", "")
	}

	dockerImageVersion := domain.DockerImageVersions{Name: languageName, Versions: []string{}}
	var dockerImageDetailSortedByVersion = make(map[string][]domain.ImageNameDetail)

	for _, imageTags := range allImageForOnelanguage.Tags {
		imageDetail, _ := splitImageDockerService.SplitDockerImageName(languageName+ ":" +imageTags)
		dockerImageDetailSortedByVersion[imageDetail.Version] = append(dockerImageDetailSortedByVersion[imageDetail.Version], imageDetail)

		if !sliceUtils.StringInSlice(imageDetail.Version, dockerImageVersion.Versions) {
			dockerImageVersion.Versions = append(dockerImageVersion.Versions, imageDetail.Version)
		}
	}

	for version, ImageNameDetailList := range dockerImageDetailSortedByVersion {
		cacheKey := "tags_" + languageName + ":" + version
		go srv.redisRepository.Add(cacheKey, ImageNameDetailList)
	}

	return dockerImageVersion, err
}

func (srv *imageDockerService) GetAllTagsFromImageVersion(languageName string, version string) ([]domain.ImageNameDetail, error) {

	cacheKey := "tags_" + languageName + ":" + version

	response := srv.redisRepository.FindDockerImageResult(cacheKey)
	fmt.Println(response[0])
	fmt.Println(len(response) > 1)

	var allImageTagsDetail []domain.ImageNameDetail

	if len(response) > 0  {
		err := json.Unmarshal([]byte(response[0]), &allImageTagsDetail)
		if err != nil {
			return nil,  errors.New(apperrors.Internal, err, "An internal error occured while searching the tags", "")
		}
	} else {
		allImageTagsDetail = srv.dbRepository.FindAllPortForLanguageAndVersion(languageName, version)
	}

	return allImageTagsDetail, nil
}
