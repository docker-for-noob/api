package splitImageDockerService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"strings"
)

func SplitDockerImageName(imageName string) (domain.ImageNameDetail, error){

	splitedLanguangeAndTags := strings.Split(imageName, ":")

	language := splitedLanguangeAndTags[0]
	allTags := splitedLanguangeAndTags[1]

	splitedTags := strings.Split(allTags, "-")

	version := splitedTags[0]

	tags := []string{}
	if len(splitedTags) > 1 {
		for _, tag := range splitedTags[1:]{
			tags = append(tags, tag)
		}
	}

	slicedImageName := domain.ImageNameDetail{
		Name:     imageName,
		Language: language,
		Version:  version,
		Tags:     tags,
	}

	return slicedImageName, nil
}