package repositories

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/JwtHelpers"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-chi/jwtauth/v5"
	"github.com/matiasvarela/errors"
	"time"
)

type JWTRepository struct{}

func NewJWTRepository() *JWTRepository {
	return &JWTRepository{}
}

var tokenAuth *jwtauth.JWTAuth

func (m JWTRepository) CreateJWTTokenString(user domain.User) (domain.JwtToken, error) {

	csrfSecret := goDotEnv.GetEnvVariable("CSRF_SECRET")

	tokenAuth = JwtHelpers.GetTokenAuth()

	expirationTime := time.Now().Add(10 * time.Minute).Unix()
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"user_id":   user.ID,
		"CSRF":      csrfSecret,
		"NotBefore": time.Now().Unix(),
		"ExpiresAt": expirationTime,
		"IssuedAt":  time.Now().Unix(),
	})
	token := domain.JwtToken{
		Data: tokenString,
	}

	if err != nil {
		return domain.JwtToken{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return token, nil
}
