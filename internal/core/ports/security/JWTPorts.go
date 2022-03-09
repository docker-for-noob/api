package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type JWTRepository interface {
	CreateJWTTokenString(user domain.User) (domain.JwtToken, error)
}
