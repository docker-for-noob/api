package dockerHubRepository

import (
	"fmt"
	domain "github.com/docker-generator/api/internal/core/domain/dockerHubDomain"
	"github.com/m7shapan/njson"
	"io/ioutil"
	"log"
	"net/http"
)

type dockerHubApi struct{}

func New() *dockerHubApi {
	return &dockerHubApi{}
}

func (repo *dockerHubApi) Read(image string, tag string) (domain.DockerHubResult, error) {
	resp, err := http.Get("https://hub.docker.com/v2/repositories/library/" + image + "/tags/?name=" + tag)
	if err != nil {
		log.Fatal(err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	var dockerHubImage domain.DockerHubImage

	err = njson.Unmarshal(jsonDataFromHttp, &dockerHubImage)
	if err != nil {
		fmt.Print(err)
	}

	var dockerHubTags []domain.DockerHubTags

	for _, data := range dockerHubImage.Results {
		var finalData domain.DockerHubTags
		errormessage := njson.Unmarshal([]byte(string(data)), &finalData)
		if errormessage != nil {
			fmt.Println(errormessage)
		}
		dockerHubTags = append(dockerHubTags, finalData)
	}

	dockerHubResult := domain.DockerHubResult{
		Name: image,
		Tags: dockerHubTags,
	}

	fmt.Print(dockerHubResult)

	return dockerHubResult, nil
}

func (repo *dockerHubApi) ReadAll() (*http.Response, error) {
	resp, err := http.Get("https://hub.docker.com/v2/repositories/library")
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}
