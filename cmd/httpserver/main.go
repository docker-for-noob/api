package main

import (
	"fmt"
	_ "github.com/docker-generator/api/docs"
	"github.com/docker-generator/api/internal/core/services/dockerHubService"
	"github.com/docker-generator/api/internal/handlers/dockerHubHandlers"
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
	dockerHubRepositoryInstanciated := repositories.NewDockerHubApi()

	dockerHubServiceInstanciated := dockerHubService.New(dockerHubRepositoryInstanciated)

	dockerHubHandler := dockerHubHandlers.NewHTTPHandler(dockerHubServiceInstanciated)

	referenceHandler := imageReferenceHandler.NewImageReferenceHandler()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://app.hetic.camillearsac.fr"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	router.Group(func(publicRouter chi.Router) {
		publicRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("toto")
		})
		publicRouter.Get("/dockerHub/images/{image}/*", dockerHubHandler.Get)
		publicRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(":8080/swagger/doc.json"), //The url pointing to API definition
		))
		publicRouter.Get("/reference/{image}", referenceHandler.Get)
	})

	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
}
