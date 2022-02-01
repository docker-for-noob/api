package dockerHubHandler

import (
	"fmt"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HTTPHandler struct {
	dockerHubService ports.DockerHubService
}

func NewHTTPHandler(dockerHubService ports.DockerHubService) *HTTPHandler {
	return &HTTPHandler{
		dockerHubService: dockerHubService,
	}
}

func (h HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.dockerHubService.GetAll()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(responseData)
}

func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	image := chi.URLParam(r, "image")
	tag := chi.URLParam(r, "*")

	resp, err := h.dockerHubService.Get(image, tag)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, resp)
	return
}
