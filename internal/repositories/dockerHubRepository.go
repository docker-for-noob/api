package repositories

import (
	"context"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/go-redis/redis/v8"
	"github.com/m7shapan/njson"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type dockerHubRepository struct {
	ctx context.Context
	rdb *redis.Client
}

func NewDockerHubRepository() *dockerHubRepository {
	ctx := context.Background()
	rdb, _ := GetRedisClient(ctx)
	return &dockerHubRepository{rdb: rdb, ctx: ctx}
}

func GetDockerHubResult(image string, tag string, page int) (domain.DockerHubImage, error) {
	var dockerHubImage domain.DockerHubImage

	resp, err := http.Get("https://hub.docker.com/v2/repositories/library/" + image + "/tags/?name=" + tag + "&page=" + strconv.Itoa(page) + "&page_size=100")

	if err != nil {
		log.Fatal(err)
	}

	jsonDataFromHttp, err := io.ReadAll(resp.Body)

	err = njson.Unmarshal(jsonDataFromHttp, &dockerHubImage)
	return dockerHubImage, err
}

func (repo *dockerHubRepository) Read(image string, tag string) (domain.DockerImageResult, error) {
	var dockerHubResult []string
	var dockerHubTags []string
	var page = 1

	for true {
		resp, err := GetDockerHubResult(image, tag, page)

		if err != nil {
			DockerImageResult := domain.DockerImageResult{}
			return DockerImageResult, nil
		}

		dockerHubResult = append(dockerHubResult, resp.Results...)

		if resp.Next == "" {
			break
		}

		page += 1
	}

	for _, data := range dockerHubResult {
		var finalData domain.DockerHubTags
		errormessage := njson.Unmarshal([]byte(data), &finalData)

		if errormessage != nil {
			log.Fatal(errormessage)
		}

		encoded, _ := json.Marshal(finalData.Tag)

		repo.rdb.RPush(repo.ctx, image+"-"+tag, strings.Replace(string(encoded), "\"", "", -1))

		dockerHubTags = append(dockerHubTags, strings.Replace(string(encoded), "\"", "", -1))
	}

	if len(dockerHubTags) == 0 {
		DockerImageResult := domain.DockerImageResult{
			Name: image,
			Tags: dockerHubTags,
		}
		return DockerImageResult, nil
	}

	DockerImageResult := domain.DockerImageResult{
		Name: image,
		Tags: dockerHubTags,
	}
	return DockerImageResult, nil
}
