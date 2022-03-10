package handlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/matiasvarela/errors"
	"net/http"
	"strconv"
)

type dockerComposeHTTPHandler struct {
	service ports.DockerComposeService
}

func NewDockerComposeHTTPHandler(dockerComposeService ports.DockerComposeService) *dockerComposeHTTPHandler {
	return &dockerComposeHTTPHandler{
		service: dockerComposeService,
	}
}

// SaveDockerCompose
// @Summary Save the json needed for docker-compose.yml file
// @Param token  header  string  true  "Bearer Token"
// @Tags Docker Compose
// @Accept  json
// @Produce json
// @Param docker-compose  body  domain.DockerCompose  true  "docker compose field"
// @Success      201  {object}  object
// @Failure      404  {object}  object
// @Router /docker-compose [post]
func (handler dockerComposeHTTPHandler) SaveDockerCompose(w http.ResponseWriter, r *http.Request) {

	dockerCompose := &domain.DockerCompose{}
	err := json.NewDecoder(r.Body).Decode(dockerCompose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, claims, _ := jwtauth.FromContext(r.Context())

	_, err = handler.service.Post(*dockerCompose, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusCreated)
}

// UpdateDockerCompose
// @Summary update the json needed for docker-compose.yml file
// @Param token  header  string  true  "Bearer Token"
// @Tags Docker Compose
// @Accept  json
// @Produce json
// @Param docker-compose  body  domain.DockerCompose  true  "docker compose field"
// @Success      204  {object}  object
// @Failure      404  {object}  object
// @Router /docker-compose [put]
func (handler dockerComposeHTTPHandler) UpdateDockerCompose(w http.ResponseWriter, r *http.Request) {
	dockerCompose := &domain.DockerCompose{}
	err := json.NewDecoder(r.Body).Decode(dockerCompose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	_, err = handler.service.Patch(*dockerCompose, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusNoContent)
}

// DeleteDockerCompose
// @Summary delete the json needed for docker-compose.yml file
// @Tags Docker Compose
// @Param token  header  string  true  "Bearer Token"
// @Param id  path  string  true  "ID"
// @Accept  json
// @Produce json
// @Success      204  {object}  object
// @Failure      404  {object}  object
// @Router /docker-compose/{id} [delete]
func (handler dockerComposeHTTPHandler) DeleteDockerCompose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, claims, _ := jwtauth.FromContext(r.Context())
	_, err := handler.service.Delete(id, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusNoContent)
}

// FindOneDockerCompose
// @Summary fetch one json needed for docker-compose.yml file
// @Tags Docker Compose
// @Param token  header  string  true  "Bearer Token"
// @Param id  path  string  true  "ID"
// @Accept  json
// @Produce json
// @Success      200  {object}  domain.DockerCompose
// @Failure      404  {object}  object
// @Router /docker-compose/{id} [get]
func (handler dockerComposeHTTPHandler) FindOneDockerCompose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, claims, _ := jwtauth.FromContext(r.Context())
	dockerCompose, err := handler.service.Get(id, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(dockerCompose)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// FindAllDockerCompose
// @Summary fetch all json needed for docker-compose.yml file
// @Tags Docker Compose
// @Param token  header  string  true  "Bearer Token"
// @Param fromItem  path  string  true  "from Item"
// @Accept  json
// @Produce json
// @Success      200  {object}  domain.AllDockerComposeDTO
// @Failure      404  {object}  object
// @Router /docker-compose/get-all/{fromItem} [get]
func (handler dockerComposeHTTPHandler) FindAllDockerCompose(w http.ResponseWriter, r *http.Request) {
	fromItemString := chi.URLParam(r, "fromItem")
	fromItem, err := strconv.Atoi(fromItemString)

	_, claims, _ := jwtauth.FromContext(r.Context())
	lastItem, dockerComposeList, err := handler.service.GetAll(fromItem, claims["user_id"].(string))

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultDatas := domain.AllDockerComposeDTO{
		LastItem: lastItem,
		Datas:    dockerComposeList,
	}

	result, err := json.Marshal(resultDatas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
