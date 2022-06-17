package main

import (
	"fmt"
	_ "github.com/docker-generator/api/docs"
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
	"github.com/docker-generator/api/pkg/uidgen"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title           Docker Generator Api
// @version         1.0.0.0
// @description     Base Go Api for Docker Generator

// @securityDefinitions.basic  JWT-Auth
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
		uidgen.New(),
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
		uidgen.New(),
	)
	dockerHubServiceInstanciated := dockerHubService.New(dockerHubRepositoryInstanciated)

	dockerComposeHandler := handlers.NewDockerComposeHTTPHandler(dockerComposeServiceInstanciated)
	versionHandler := handlers.NewVersionHTTPHandler(versionServiceInstanciated)
	userHandler := userHandlers.New(userServiceInstanciated)
	authentificationHandler := authentificationHandlers.New(authentificationwithJWTServiceInstanciated)
	dockerHubHandler := dockerHubHandlers.NewHTTPHandler(dockerHubServiceInstanciated)

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
		publicRouter.Post("/user", userHandler.Post)
		publicRouter.Post("/authentication/login", authentificationHandler.Login)
		publicRouter.Get("/dockerHub/images/{image}/*", dockerHubHandler.Get)
		publicRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(":8080/swagger/doc.json"), //The url pointing to API definition
		))
	})

	router.Group(func(privateRouter chi.Router) {
		privateRouter.Use(jwtauth.Verifier(JwtHelpers.GetTokenAuth()))
		privateRouter.Use(jwtauth.Authenticator)

		privateRouter.Post("/authentication/logout", authentificationHandler.Logout)
		privateRouter.Post("/docker-compose", dockerComposeHandler.SaveDockerCompose)
		privateRouter.Put("/docker-compose", dockerComposeHandler.UpdateDockerCompose)
		privateRouter.Delete("/docker-compose/{id}", dockerComposeHandler.DeleteDockerCompose)
		privateRouter.Get("/docker-compose/{id}", dockerComposeHandler.FindOneDockerCompose)
		privateRouter.Get("/docker-compose/get-all/{fromItem}", dockerComposeHandler.FindAllDockerCompose)
		privateRouter.Get("/docker-compose/{id}/version/{versionId}", versionHandler.FindOneVersion)
		privateRouter.Get("/docker-compose/{id}/version", versionHandler.FindAllVersion)
		privateRouter.Get("/user", userHandler.Get)
		privateRouter.Patch("/user", userHandler.Patch)
		privateRouter.Delete("/user", userHandler.Delete)
	})

	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
}
