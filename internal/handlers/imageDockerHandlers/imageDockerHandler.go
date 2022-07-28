package imageDockerHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HTTPHandler struct {
	imageDockerService ports.ImageDockerService
}

func NewHTTPHandler(imageDockerService ports.ImageDockerService) *HTTPHandler {
	return &HTTPHandler{
		imageDockerService: imageDockerService,
	}
}

// Get
// @Summary  returns a docker image from docker hub or redis
// @Tags Docker Hub
// @Param image   path      string  true  "Docker hub image"
// @Param tag   path      string  false  "Docker hub tag"
// @Accept  json
// @Produce json
// @Success      200  {object}  domain.DockerImageResult
// @Failure      404  {object}  object
// @Router /dockerHub/images/{image}/{tag} [get]
func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	image := chi.URLParam(r, "image")
	tag := chi.URLParam(r, "*")

	resp, err := h.imageDockerService.Get(image, tag)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if resp.Name == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Image doesn't exist"))
		return
	}

	if resp.Name != "" && resp.Tags == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Tags doesn't exist"))
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
