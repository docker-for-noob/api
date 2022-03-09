package authentificationWithJWTService

import (
	"github.com/docker-generator/api/internal/core/domain"
	authentificationPorts "github.com/docker-generator/api/internal/core/ports/authentification"
	securityPorts "github.com/docker-generator/api/internal/core/ports/security"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
)

type authentificationWithCookieService struct {
	authentificationRepository authentificationPorts.AuthentificationRepository
	JWTRepository              securityPorts.JWTRepository
}

func New(
	authentificationRepository authentificationPorts.AuthentificationRepository,
	JWTRepository securityPorts.JWTRepository,
) *authentificationWithCookieService {
	return &authentificationWithCookieService{
		authentificationRepository: authentificationRepository,
		JWTRepository:              JWTRepository,
	}
}

func (a authentificationWithCookieService) Login(creds domain.Credentials) (domain.JwtToken, error) {
	user, err := a.authentificationRepository.Login(creds)
	if err != nil {
		return domain.JwtToken{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	token, err := a.JWTRepository.CreateJWTTokenString(user)

	if err != nil {
		return domain.JwtToken{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return token, err
}

func (a authentificationWithCookieService) Logout(id string) error {
	err := a.authentificationRepository.Logout(id)

	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return nil
}
