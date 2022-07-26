package main

import (
	"fmt"
	_ "github.com/docker-generator/api/docs"
	"github.com/docker-generator/api/internal/core/services/imageDockerService"
	"github.com/docker-generator/api/internal/core/services/imageReferenceService"
	"github.com/docker-generator/api/internal/handlers/imageDockerHandlers"
	"github.com/docker-generator/api/internal/handlers/imageReferenceHandler"
	"github.com/docker-generator/api/internal/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title           Docker Generator Api
// @version         1.0.0.0
// @description     Base Go Api for Docker Generator

// @securityDefinitions.basic  JWT-Auth
func main() {
	dockerHubRepositoryInstantiated := repositories.NewDockerHubRepository()

	redisRepositoryInstantiated := repositories.NewRedisRepository()

	imageDockerServiceInstantiated := imageDockerService.New(dockerHubRepositoryInstantiated, redisRepositoryInstantiated)

	imageDockerHandler := imageDockerHandlers.NewHTTPHandler(imageDockerServiceInstantiated)

	imageReferenceRepositoryInstantiated := repositories.NewImageReferenceRepository()

	imageReferenceServiceInstantiated := imageReferenceService.New(imageReferenceRepositoryInstantiated, dockerHubRepositoryInstantiated)

	referenceHandler := imageReferenceHandler.NewImageReferenceHandler(imageReferenceServiceInstantiated)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	router.Group(func(publicRouter chi.Router) {
		publicRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("toto")
		})
		publicRouter.Get("/dockerHub/images/{image}/*", imageDockerHandler.Get)
		publicRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(":80/swagger/doc.json"), //The url pointing to API definition
		))
		publicRouter.Get("/reference/{image}", referenceHandler.Get)
	})

	err := http.ListenAndServe(":80", router)
	fmt.Println(err)
}
