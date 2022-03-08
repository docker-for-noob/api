package handlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/go-chi/chi/v5"
	"github.com/matiasvarela/errors"
	"net/http"
	"strconv"
)

type dockerComposeHTTPHandler struct {
	service ports.DockerComposeService
}

func NewDockerComposeHTTPHandler(dockerComposeService ports.DockerComposeService) *dockerComposeHTTPHandler{
	return &dockerComposeHTTPHandler{
		service: dockerComposeService,
	}
}

func (handler dockerComposeHTTPHandler) SaveDockerCompose(w http.ResponseWriter, r *http.Request){

	dockerCompose := &domain.DockerCompose{}
	err := json.NewDecoder(r.Body).Decode(dockerCompose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = handler.service.Post(*dockerCompose)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusCreated)
}

func (handler dockerComposeHTTPHandler) UpdateDockerCompose(w http.ResponseWriter, r *http.Request) {
	dockerCompose := &domain.DockerCompose{}
	err := json.NewDecoder(r.Body).Decode(dockerCompose)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = handler.service.Patch(*dockerCompose)
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

func (handler dockerComposeHTTPHandler) DeleteDockerCompose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := handler.service.Delete(id)
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

func (handler dockerComposeHTTPHandler) FindOneDockerCompose(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	dockerCompose, err := handler.service.Get(id)
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

func (handler dockerComposeHTTPHandler) FindAllDockerCompose(w http.ResponseWriter, r *http.Request) {
	fromItemString := chi.URLParam(r, "fromItem")
	fromItem, err :=strconv.Atoi(fromItemString)

	lastItem, dockerComposeList, err := handler.service.GetAll(fromItem)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultDatas := struct {
		LastItem int
		Datas []domain.DockerCompose
	}{
		lastItem,
		dockerComposeList,
	}

	result, err := json.Marshal(resultDatas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


