package userService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/ports/user"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/docker-generator/api/pkg/formater"
	"github.com/docker-generator/api/pkg/uidgen"
	"github.com/matiasvarela/errors"
)

type userService struct {
	userRepository              ports.UserRepository
	BCryptRepository            ports.BCryptRepository
	passwordValidatorRepository ports.PasswordValidatorRepository
	uuidgenerator               uidgen.UIDGen
}

func New(
	userRepository ports.UserRepository,
	BCryptRepository ports.BCryptRepository,
	passwordValidatorRepository ports.PasswordValidatorRepository,
	uuidgenerator uidgen.UIDGen,
) *userService {
	return &userService{
		userRepository:              userRepository,
		BCryptRepository:            BCryptRepository,
		passwordValidatorRepository: passwordValidatorRepository,
		uuidgenerator:               uuidgenerator,
	}
}

func (srv *userService) Get(id string) (domain.User, error) {
	user, err := srv.userRepository.Read(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.User{}, errors.New(apperrors.NotFound, err, "User not found in database", "")
		}

		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return user, nil
}

func (srv *userService) Post(user domain.User) (domain.User, error) {

	err := srv.passwordValidatorRepository.VerifyPasswordStrenght(user.Password)
	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	pass, err := srv.BCryptRepository.HashPassword(user.Password)

	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	user.ID = srv.uuidgenerator.NewUuid()
	user.Password = pass

	user.Email = formater.NormalizeEmail(user.Email)

	_, err = srv.userRepository.Create(user)

	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return user, nil
}
func (srv *userService) Patch(id string, user domain.User) (domain.User, error) {
	user, err := srv.userRepository.Update(id, user)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.User{}, errors.New(apperrors.NotFound, err, "User not found in database", "")
		}
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return user, nil
}
func (srv *userService) Delete(id string) (bool, error) {
	isDeleted, err := srv.userRepository.Delete(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return isDeleted, errors.New(apperrors.NotFound, err, "User not found in database", "")
		}
		return isDeleted, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return isDeleted, nil
}
