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

type dockerHubApi struct{}

func NewDockerHubApi() *dockerHubApi {
	return &dockerHubApi{}
}

func (repo *dockerHubApi) Read(image string, tag string) (domain.DockerHubResult, error) {
	ctx := context.Background()

	rdb := GetClient(ctx)
	length := rdb.LLen(ctx, image+"-"+tag).Val()

	var dockerHubTags []string

	if length > 0 {
		dockerHubTags = rdb.LRange(ctx, image+"-"+tag, 0, -1).Val()
	} else {
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

		for _, data := range dockerHubImage.Results {
			var finalData domain.DockerHubTags
			errormessage := njson.Unmarshal([]byte(data), &finalData)

			if errormessage != nil {
				log.Fatal(err)
			}

			encoded, _ := json.Marshal(finalData.Tag)

			rdb.RPush(ctx, image+"-"+tag, encoded)

			dockerHubTags = append(dockerHubTags, string(encoded))
		}
	}

	dockerHubResult := domain.DockerHubResult{
		Name: image,
		Tags: dockerHubTags,
	}
	return dockerHubResult, nil
}
