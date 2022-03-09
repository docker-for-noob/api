package ports

import (
	"github.com/docker-generator/api/internal/core/domain"
)

type UserRepository interface {
	Read(id string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(id string, user domain.User) (domain.User, error)
	Delete(id string) (bool, error)
}

type UserService interface {
	Get(id string) (domain.User, error)
	Post(user domain.User) (domain.User, error)
	Patch(id string, user domain.User) (domain.User, error)
	Delete(id string) (bool, error)
}
