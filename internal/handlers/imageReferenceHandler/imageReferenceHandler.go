package imageReferenceHandler

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type imageReferenceHandler struct {}

func NewImageReferenceHandler() *imageReferenceHandler {
	return &imageReferenceHandler{}
}

// Get
// @Summary  returns a docker image from docker hub or redis
// @Tags Docker Hub
// @Param image   path      string  true  "Docker hub image"
// @Param tag   path      string  false  "Docker hub tag"
// @Accept  json
// @Produce json
// @Success      200  {object}  domain.DockerHubResult
// @Failure      404  {object}  object
// @Router /dockerHub/images/{image}/{tag} [get]
func (h imageReferenceHandler) Get(w http.ResponseWriter, r *http.Request) {
	image := chi.URLParam(r, "image")


	if len(image) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := domain.ImageReference{}

	respEnv := domain.EnvVar{}

	respEnv.Key = "MYSQL_ROOT_PASSWORD"
	respEnv.Desc = "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw."

	resp.Id = uuid.New()
	resp.Name = "go:latest"
	resp.Port = []string{"8080"}
	resp.Workdir = []string{"path/to/file"}
	resp.Env = []domain.EnvVar{respEnv}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	result, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(result)
}