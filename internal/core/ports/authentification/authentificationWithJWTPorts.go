package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type AuthentificationWithJWTService interface {
	Login(credentials domain.Credentials) (domain.JwtToken, error)
	Logout(id string) error
}
