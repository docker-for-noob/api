package authentificationWithJWTService

import (
	"github.com/docker-generator/api/internal/core/domain"
	authentificationPorts "github.com/docker-generator/api/internal/core/ports/authentification"
	securityPorts "github.com/docker-generator/api/internal/core/ports/security"
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

	cookie, _ := a.JWTRepository.CreateJWTTokenString(user)

	return cookie, err
}

func (a authentificationWithCookieService) Logout(id string) error {

	return a.authentificationRepository.Logout(id)
}
