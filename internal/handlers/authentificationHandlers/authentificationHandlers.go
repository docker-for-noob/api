package authentificationHandlers

import (
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	ports "github.com/docker-generator/api/internal/core/ports/authentification"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"time"
)

type HTTPHandler struct {
	authentificationWithJWTService ports.AuthentificationWithJWTService
}

func New(authentificationWithJWTService ports.AuthentificationWithJWTService) *HTTPHandler {
	return &HTTPHandler{
		authentificationWithJWTService: authentificationWithJWTService,
	}
}

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

	cookie := http.Cookie{
		Name:     "jwt",
		HttpOnly: true,
		Value:    token.Data,
		Path:     "/",
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	http.SetCookie(w, &cookie)
}

func (h HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	err := h.authentificationWithJWTService.Logout(claims["user_id"].(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:    "jwt",
		MaxAge:  -1,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	http.SetCookie(w, &cookie)
}
