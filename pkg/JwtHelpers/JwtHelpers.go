package JwtHelpers

import (
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	jwtSecret := goDotEnv.GetEnvVariable("JWT_SECRET")
	tokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}
