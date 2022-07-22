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

func (repo *dockerHubRepository) Read(image string, tag string) (domain.DockerImageResult, error) {
	var dockerHubTags []string

	// TODO : Boucle sur toutes les pages
	resp, err := http.Get("https://hub.docker.com/v2/repositories/library/" + image + "/tags/?name=" + tag + "&page_size=100")

	if resp.StatusCode == 403 {
		DockerImageResult := domain.DockerImageResult{}
		return DockerImageResult, nil
	}

	if err != nil {
		log.Fatal(err)
	}

	jsonDataFromHttp, err := io.ReadAll(resp.Body)

	var dockerHubImage domain.DockerHubImage

	err = njson.Unmarshal(jsonDataFromHttp, &dockerHubImage)

	if err != nil {
		log.Fatal(err)
	}

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
