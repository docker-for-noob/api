package repositories

import (
	"context"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/m7shapan/njson"
	"io/ioutil"
	"log"
	"net/http"
)

type dockerHubRepository struct{}

func NewDockerHubRepository() *dockerHubRepository {
	return &dockerHubRepository{}
}

func (repo *dockerHubRepository) Read(image string, tag string) (domain.DockerImageResult, error) {
	var dockerHubTags []string

	resp, err := http.Get("https://hub.docker.com/v2/repositories/library/" + image + "/tags/?name=" + tag)

	if err != nil {
		log.Fatal(err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	var dockerHubImage domain.DockerHubImage

	err = njson.Unmarshal(jsonDataFromHttp, &dockerHubImage)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	redisRepository := NewRedisRepository()
	_, err = redisRepository.GetRedisClient(ctx)

	if err != nil {
		return domain.DockerImageResult{}, err
	}

	for _, data := range dockerHubImage.Results {
		var finalData domain.DockerHubTags
		errormessage := njson.Unmarshal([]byte(data), &finalData)

		if errormessage != nil {
			log.Fatal(err)
		}

		encoded, _ := json.Marshal(finalData.Tag)

		redisRepository.AddImage(image, tag, encoded)

		dockerHubTags = append(dockerHubTags, string(encoded))
	}

	DockerImageResult := domain.DockerImageResult{
		Name: image,
		Tags: dockerHubTags,
	}
	return DockerImageResult, nil
}
