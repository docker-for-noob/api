package dockerHubRepository

import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/docker-generator/api/internal/core/domain/dockerHubDomain"
	"github.com/docker-generator/api/internal/repositories/redisRepository"
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
	ctx := context.Background()

	rdb, _ := redisRepository.GetClient(ctx)
	length := rdb.LLen(ctx, image+"-"+tag).Val()

	var dockerHubTags []string

	if length > 0 {
		fmt.Println("------------- REDIS -----------------")
		fmt.Println(rdb.LRange(ctx, image+"-"+tag, 0, -1).Val())

		dockerHubTags = rdb.LRange(ctx, image+"-"+tag, 0, -1).Val()
	} else {
		fmt.Println("-------------- DOCKER HUB ----------------")
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

		for _, data := range dockerHubImage.Results {
			var finalData domain.DockerHubTags
			errormessage := njson.Unmarshal([]byte(data), &finalData)

			if errormessage != nil {
				fmt.Println(errormessage)
			}

			encoded, _ := json.Marshal(finalData.Tag)

			rdb.RPush(ctx, image+"-"+tag, encoded)

			dockerHubTags = append(dockerHubTags, string(encoded))
		}
		fmt.Println(rdb.LRange(ctx, image+"-"+tag, 0, -1).Val())
	}

	dockerHubResult := domain.DockerHubResult{
		Name: image,
		Tags: dockerHubTags,
	}
	return dockerHubResult, nil
}

func (repo *dockerHubApi) ReadAll() (*http.Response, error) {
	resp, err := http.Get("https://hub.docker.com/v2/repositories/library")
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}
