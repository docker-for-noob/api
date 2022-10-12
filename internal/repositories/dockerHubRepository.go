package repositories

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-redis/redis/v8"
	"github.com/m7shapan/njson"
	"github.com/matiasvarela/errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func GetDockerHubOfficialImageResult(page int) (domain.DockerImages, error) {
	var dockerImages domain.DockerImages

	resp, err := http.Get("https://hub.docker.com/v2/repositories/library?page=" + strconv.Itoa(page) + "&page_size=100")

	if err != nil {
		log.Fatal(err)
	}

	jsonDataFromHttp, err := io.ReadAll(resp.Body)

	err = njson.Unmarshal(jsonDataFromHttp, &dockerImages)

	return dockerImages, err
}

func (repo *dockerHubRepository) GetImages() (domain.DockerImagesParse, error) {
	var dockerImagesResults []string
	var page = 1

	for true {
		resp, _ := GetDockerHubOfficialImageResult(page)

		for _, data := range resp.Results {
			dockerImagesResults = append(dockerImagesResults, data.Name)
		}

		if resp.Next == "" {
			break
		}

		page += 1
	}

	dockerImagesParse := domain.DockerImagesParse{
		Images: dockerImagesResults,
	}

	return dockerImagesParse, nil
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

func (repo *dockerHubRepository) GetTagReference(image string, tag string) (domain.ImageReference, error) {

	resp, err := http.Get("https://hub.docker.com/v2/repositories/library/" + image + "/tags/" + tag + "/images")
	if err != nil {
		return domain.ImageReference{}, err
	}

	remainingRequest := resp.Header.Get("x-ratelimit-remaining")
	timeStampUntilNextRequest := resp.Header.Get("x-ratelimit-reset")

	dataFromHttp, _ := io.ReadAll(resp.Body)

	var instruction []struct {
		Layers []struct {
			Instruction string `json:"instruction"`
		} `json:"layers"`
	}

	errormessage := json.Unmarshal(dataFromHttp, &instruction)
	if errormessage != nil {
		return domain.ImageReference{}, errors.New(apperrors.Internal, errormessage, "An internal error occured while searching the reference", "")
	}

	var imageReference = domain.ImageReference{}
	imageReference.Name = image + ":" + tag

	for _, data := range instruction[0].Layers {
		words := strings.Fields(data.Instruction)
		if words[0] == "VOLUME" || words[0] == "WORKDIR" {
			for i := 1; i < len(words); i++ {
				imageReference.Workdir = append(imageReference.Workdir, words[i])
			}
		}
		if words[0] == "EXPOSE" {
			for i := 1; i < len(words); i++ {
				imageReference.Port = append(imageReference.Port, words[i])
			}
		}
	}

	if remainingRequest == "0" {
		responseTimeStampInInt, _ := strconv.ParseInt(timeStampUntilNextRequest, 10, 64)
		shouldIWait := time.Now().Unix() < responseTimeStampInInt
		for shouldIWait {
			time.Sleep(60 * time.Second)
			shouldIWait = time.Now().Unix() < responseTimeStampInInt
		}
	}

	return imageReference, nil
}

func (repo *dockerHubRepository) HandleMultipleGetTagReference(image string, allTag []string) error {

	pathToBuffer := goDotEnv.GetEnvVariable("BATCH_REFERENTIEL_BUFFER")
	if _, err := os.Stat(pathToBuffer); err != nil {
		return err
	}

	f, err := os.OpenFile(pathToBuffer, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(f)

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	for _, tagName := range allTag {
		tagReference, err := repo.GetTagReference(image, tagName)
		if err != nil {
			time.Sleep(60 * time.Second)
			tagReference, err = repo.GetTagReference(image, tagName)
			if err != nil {
				return err
			}
		}
		if len(tagReference.Port) > 0 && len(tagReference.Workdir) > 0 {
			values := tagReferenceToSlice(tagReference)
			err = csvwriter.Write(values)
			if err != nil {
				return err
			}
			csvwriter.Flush()
		}
	}

	return nil
}

func tagReferenceToSlice(reference domain.ImageReference) []string {

	var portsToString string
	var workDirToSting string
	var envToSting string

	for _, element := range reference.Port {
		if len(portsToString) == 0 {
			portsToString = element
		} else {
			portsToString = portsToString + " " + element
		}

	}

	for _, element := range reference.Workdir {
		element = strings.ReplaceAll(element, "[", "")
		element = strings.ReplaceAll(element, "]", "")
		if len(workDirToSting) == 0 {
			workDirToSting = element
		} else {
			workDirToSting = workDirToSting + " " + element
		}
	}

	for _, element := range reference.Env {
		if len(envToSting) == 0 {
			envToSting = element.Key + "=" + element.Desc
		} else {
			envToSting = envToSting + " " + element.Key + "=" + element.Desc
		}
	}

	slice := []string{
		reference.Name,
		portsToString,
		workDirToSting,
		envToSting,
	}

	return slice
}
