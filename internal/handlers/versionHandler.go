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
)

type versionHTTPHandler struct {
	service ports.VersionService
}

func NewVersionHTTPHandler(versionService ports.VersionService) *versionHTTPHandler {
	return &versionHTTPHandler{
		service: versionService,
	}
}

func (handler versionHTTPHandler) FindOneVersion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	versionId := chi.URLParam(r, "versionId")
	_, claims, _ := jwtauth.FromContext(r.Context())

	dockerComposeVersion, err := handler.service.Get(id, versionId, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(dockerComposeVersion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (handler versionHTTPHandler) FindAllVersion(w http.ResponseWriter, r *http.Request) {
	idDockerCompose := chi.URLParam(r, "id")
	_, claims, _ := jwtauth.FromContext(r.Context())

	allVersions, err := handler.service.GetAll(idDockerCompose, claims["user_id"].(string))
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultDatas := struct {
		Datas []domain.DockerCompose
	}{
		allVersions,
	}

	result, err := json.Marshal(resultDatas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
