package main

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/services/authentification/authentificationWithJWTService"
	"github.com/docker-generator/api/internal/core/services/dockerComposeService"
	"github.com/docker-generator/api/internal/core/services/dockerHubService"
	"github.com/docker-generator/api/internal/core/services/userService"
	"github.com/docker-generator/api/internal/core/services/versionService"
	"github.com/docker-generator/api/internal/handlers"
	"github.com/docker-generator/api/internal/handlers/authentificationHandlers"
	"github.com/docker-generator/api/internal/handlers/dockerHubHandlers"
	"github.com/docker-generator/api/internal/handlers/userHandlers"
	"github.com/docker-generator/api/internal/repositories"
	"github.com/docker-generator/api/pkg/JwtHelpers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func main() {
	BcryptRepositoryInstanciated := repositories.NewBcryptRepository()
	userRepositoryInstanciated := repositories.NewUserRepository()
	JWTInstanciated := repositories.NewJWTRepository()
	authentificationRepositoryInstanciated := repositories.NewAuthentificationRepository(BcryptRepositoryInstanciated)
	passwordValidatorRepositoryInstanciated := repositories.NewPasswodValidatorRepository()
	dockerComposeRepositoryInstanciated := repositories.NewDockerComposeFirestore()
	versionFirestoreRepositoryIntanciated := repositories.NewVersionFirestore()
	dockerHubRepositoryInstanciated := repositories.NewDockerHubApi()

	userServiceInstanciated := userService.New(
		userRepositoryInstanciated,
		BcryptRepositoryInstanciated,
		passwordValidatorRepositoryInstanciated,
	)
	authentificationwithJWTServiceInstanciated := authentificationWithJWTService.New(
		authentificationRepositoryInstanciated,
		JWTInstanciated,
	)
	versionServiceInstanciated := versionService.New(
		dockerComposeRepositoryInstanciated,
		versionFirestoreRepositoryIntanciated,
	)
	dockerComposeServiceInstanciated := dockerComposeService.New(
		dockerComposeRepositoryInstanciated,
		versionServiceInstanciated,
	)

	dockerHubServiceInstanciated := dockerHubService.New(dockerHubRepositoryInstanciated)

	dockerComposeHandler := handlers.NewDockerComposeHTTPHandler(dockerComposeServiceInstanciated)
	versionHandler := handlers.NewVersionHTTPHandler(versionServiceInstanciated)
	userHandler := userHandlers.New(userServiceInstanciated)
	authentificationHandler := authentificationHandlers.New(authentificationwithJWTServiceInstanciated)
	dockerHubHandler := dockerHubHandlers.NewHTTPHandler(dockerHubServiceInstanciated)

	router := chi.NewRouter()

	router.Group(func(publicRouter chi.Router) {
		publicRouter.Post("/user/create", userHandler.Post)
		publicRouter.Get("/user/{id}", userHandler.Get)
		publicRouter.Patch("/user/update/{id}", userHandler.Patch)
		publicRouter.Delete("/user/{id}", userHandler.Delete)
		publicRouter.Post("/authentication/login", authentificationHandler.Login)
		publicRouter.Get("/dockerHub/images", dockerHubHandler.GetAll)
		publicRouter.Get("/dockerHub/images/{image}/*", dockerHubHandler.Get)
	})

	router.Group(func(privateRouter chi.Router) {
		privateRouter.Use(jwtauth.Verifier(JwtHelpers.GetTokenAuth()))
		privateRouter.Use(jwtauth.Authenticator)

		privateRouter.Post("/authentication/logout", authentificationHandler.Logout)
		privateRouter.Post("/docker-compose/save", dockerComposeHandler.SaveDockerCompose)
		privateRouter.Put("/docker-compose/update", dockerComposeHandler.UpdateDockerCompose)
		privateRouter.Delete("/docker-compose/delete/{id}", dockerComposeHandler.DeleteDockerCompose)
		privateRouter.Get("/docker-compose/{id}", dockerComposeHandler.FindOneDockerCompose)
		privateRouter.Get("/docker-compose/get-all/{fromItem}", dockerComposeHandler.FindAllDockerCompose)
		privateRouter.Get("/docker-compose/{id}/version/{versionId}", versionHandler.FindOneVersion)
		privateRouter.Get("/docker-compose/{id}/version", versionHandler.FindAllVersion)
	})

	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
}
