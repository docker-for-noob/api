package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type AuthentificationRepository interface {
	Login(credentials domain.Credentials) (domain.User, error)
	Logout(id string) error
}

type AuthentificationService interface {
	Login(credentials domain.Credentials) error
	Logout(id string) error
}
