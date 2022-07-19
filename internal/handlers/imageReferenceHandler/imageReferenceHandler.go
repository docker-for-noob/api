package imageReferenceHandler

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/ports"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/go-chi/chi/v5"
	"github.com/matiasvarela/errors"
	"net/http"
)

type imageReferenceHandler struct {
	imageReferenceService ports.ImageReferenceService
}

func NewImageReferenceHandler(imageReferenceService ports.ImageReferenceService) *imageReferenceHandler {
	return &imageReferenceHandler{
		imageReferenceService: imageReferenceService,
	}
}

// Get
// @Summary  returns a reference docker image
// @Tags Docker Hub
// @Param image   path      string  true  "Docker hub image"
// @Accept  json
// @Produce json
// @Success      200  {object}  domain.ImageReference
// @Failure      404  {object}  object
// @Router /dockerHub/images/{image}/{tag} [get]
func (h imageReferenceHandler) Get(w http.ResponseWriter, r *http.Request) {
	image := chi.URLParam(r, "image")

	if len(image) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, executionError := h.imageReferenceService.Get(image)

	if executionError != nil {
		if errors.Is(executionError, apperrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
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