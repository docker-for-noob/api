package userHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	ports "github.com/docker-generator/api/internal/core/ports/user"
	"github.com/go-chi/jwtauth/v5"
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

// Get
// @Summary Fetch his own profile
// @Tags User
// @Accept  json
// @Produce json
// @Param token  header  string  true  "Bearer Token"
// @Sucess      201  {object}  domain.User
// @Failure      404  {object}  object
// @Router /user [get]
func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	resp, err := h.userService.Get(claims["user_id"].(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Write(response)
}

// Post
// @Summary Create a User
// @Tags User
// @Accept  json
// @Produce json
// @Sucess      201  {object}  object
// @Failure      404  {object}  object
// @Router /user [post]
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

// Patch
// @Summary update his own profile
//@Tags User
// @Accept  json
// @Produce json
// @Param id path string true "User ID"
// @Param token  header  string  true  "Bearer Token"
// @Sucess      201  {object}  object
// @Failure      404  {object}  object
// @Router /user [patch]
func (h HTTPHandler) Patch(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	_, claims, _ := jwtauth.FromContext(r.Context())

	user, err := h.userService.Patch(claims["user_id"].(string), user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

// Delete
// @Summary delete his own profile
//@Tags User
// @Accept  json
// @Produce json
// @Param id path string true "User ID"
// @Param token  header  string  true  "Bearer Token"
// @Sucess      201  {object}  object
// @Failure      404  {object}  object
// @Router /user [delete]
func (h HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	isDeleted, err := h.userService.Delete(claims["user_id"].(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, _ := json.Marshal(isDeleted)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Write(response)
}
