package userHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	ports "github.com/docker-generator/api/internal/core/ports/user"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HTTPHandler struct {
	userService ports.UserService
}

func New(userService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.userService.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Write(response)
}

func (h HTTPHandler) Post(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	_, err := h.userService.Post(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func (h HTTPHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := domain.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)

	user, err := h.userService.Patch(id, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	isDeleted, err := h.userService.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, _ := json.Marshal(isDeleted)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Write(response)
}
