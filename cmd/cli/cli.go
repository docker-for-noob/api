package main

import (
	"github.com/docker-generator/api/internal/core/services/imageDockerService"
	"github.com/docker-generator/api/internal/core/services/imageReferenceService"
	"github.com/docker-generator/api/internal/handlers/referenceBatchHandler"
	"github.com/docker-generator/api/internal/repositories"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	imageReferenceRepositoryInstanciated := repositories.NewImageReferenceRepository()
	dockerHubRepositoryInstantiated := repositories.NewDockerHubRepository()
	redisRepositoryInstantiated := repositories.NewRedisRepository()
	imageDockerServiceInstantiated := imageDockerService.New(dockerHubRepositoryInstantiated, redisRepositoryInstantiated)
	imageReferencialServiceinstanciated := imageReferenceService.New(imageReferenceRepositoryInstanciated, dockerHubRepositoryInstantiated, imageDockerServiceInstantiated)
	referenceBatchHandlerInstanciated := referenceBatchHandler.NewReferenceBatchHandler(imageReferencialServiceinstanciated)

	app := &cli.App{
		Name:  "getReferenceBatch",
		Usage: "find reference for image batch",
		Action: func(*cli.Context) error {
			handlerErr := referenceBatchHandlerInstanciated.Run()
			if handlerErr != nil {
				return handlerErr
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
