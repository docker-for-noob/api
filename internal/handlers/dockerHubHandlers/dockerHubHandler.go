package dockerHubHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HTTPHandler struct {
	dockerHubService ports.DockerHubService
}

func NewHTTPHandler(dockerHubService ports.DockerHubService) *HTTPHandler {
	return &HTTPHandler{
		dockerHubService: dockerHubService,
	}
}

func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	image := chi.URLParam(r, "image")
	tag := chi.URLParam(r, "*")

	resp, err := h.dockerHubService.Get(image, tag)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	result, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(result)
}
