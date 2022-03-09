package repositories

import (
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/go-passwd/validator"
	"github.com/matiasvarela/errors"
)

type PasswodValidatorRepository struct{}

func NewPasswodValidatorRepository() *PasswodValidatorRepository {
	return &PasswodValidatorRepository{}
}

func (repo *PasswodValidatorRepository) VerifyPasswordStrenght(plainPassword string) error {
	passwordValidator := validator.New(
		validator.MinLength(8, errors.New(apperrors.InvalidInput, nil, "The password is at least 8 characters long", "")),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, errors.New(apperrors.InvalidInput, nil, "The password has at least one lower letter", "")),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, errors.New(apperrors.InvalidInput, nil, "The password has at least one upper letter", "")),
		validator.ContainsAtLeast("0123456789", 1, errors.New(apperrors.InvalidInput, nil, "The password has at least one number", "")),
	)
	return passwordValidator.Validate(plainPassword)
}
