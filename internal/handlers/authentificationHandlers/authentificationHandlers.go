package authentificationHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	ports "github.com/docker-generator/api/internal/core/ports/authentification"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

type HTTPHandler struct {
	authentificationWithJWTService ports.AuthentificationWithJWTService
}

func New(authentificationWithJWTService ports.AuthentificationWithJWTService) *HTTPHandler {
	return &HTTPHandler{
		authentificationWithJWTService: authentificationWithJWTService,
	}
}

// Login
// @Summary login into application
// @Tags Authentification
// @Accept  json
// @Produce json
// @Param credentials  body  domain.Credentials  true  "Credentials"
// @Success      201  {object}  string
// @Failure      404  {object}  object
// @Router /authentication/login [post]
func (h HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds domain.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	token, err := h.authentificationWithJWTService.Login(creds)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token.Data))
}

// Logout
// @Summary logout of application
// @Tags Authentification
// @Accept  json
// @Produce json
// @Param token  header  string  true  "Bearer Token"
// @Sucess      201  {object}  object
// @Failure      404  {object}  object
// @Router /authentication/logout [post]
func (h HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	err := h.authentificationWithJWTService.Logout(claims["user_id"].(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
