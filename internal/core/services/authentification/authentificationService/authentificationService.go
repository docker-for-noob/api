package authentificationService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports/authentification"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/docker-generator/api/pkg/formater"
	"github.com/matiasvarela/errors"
)

type authentificationService struct {
	authentificationRepository ports.AuthentificationRepository
}

func New(authentificationRepository ports.AuthentificationRepository) *authentificationService {
	return &authentificationService{
		authentificationRepository: authentificationRepository,
	}
}

func (srv *authentificationService) Login(credentials domain.Credentials) (domain.User, error) {
	credentials.Email = formater.NormalizeEmail(credentials.Email)

	User, err := srv.authentificationRepository.Login(credentials)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.User{}, errors.New(apperrors.NotFound, err, "User not found in database", "")
		}
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return User, nil
}

func (srv *authentificationService) Logout(id string) error {

	err := srv.authentificationRepository.Logout(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return errors.New(apperrors.NotFound, err, "User not found in database", "")
		}
		return errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return errors.New(apperrors.Internal, err, "An internal error occurred", "")

}
