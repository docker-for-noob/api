package main

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/services/dockerComposeService"
	"github.com/docker-generator/api/internal/handlers"
	"github.com/docker-generator/api/internal/repositories"
	"github.com/go-chi/chi/v5"
	"net/http"
)


func main() {
	dockerComposeRepositoryInstanciated := repositories.NewDockerComposeFirestore()
	dockerComposeServiceInstanciated := dockerComposeService.New(dockerComposeRepositoryInstanciated)
	dockerComposeHandler := handlers.NewDockerComposeHTTPHandler(dockerComposeServiceInstanciated)


	r := chi.NewRouter()
	r.Post("/docker-compose/save", dockerComposeHandler.SaveDockerCompose)
	err := http.ListenAndServe(":8080", r)
	fmt.Println(err)
}