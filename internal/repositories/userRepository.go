package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"github.com/mitchellh/mapstructure"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Create(user domain.User) (domain.User, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	snapshot, err := client.Collection("Users").Where("Email", "==", user.Email).Snapshots(ctx).Next()

	if snapshot.Size > 0 {
		return domain.User{}, errors.New(apperrors.InvalidInput, err, "User already exists", "")
	}

	_, err = client.Collection("Users").Doc(user.ID).Set(ctx, map[string]interface{}{
		"ID":        user.ID,
		"Email":     user.Email,
		"Password":  user.Password,
		"CreatedAt": firestore.ServerTimestamp,
	})
	if err != nil {
		return user, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return user, nil
}

func (repo *userRepository) Read(id string) (domain.User, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	userSearch, err := client.Collection("Users").Doc(id).Get(ctx)
	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	userData := userSearch.Data()
	userResult := domain.User{}

	err = mapstructure.Decode(userData, &userResult)
	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return userResult, nil
}

func (repo *userRepository) Update(id string, user domain.User) (domain.User, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc(id).Set(ctx, map[string]interface{}{
		"ID":           user.ID,
		"Email       ": user.Email,
		"Password    ": user.Password,
	}, firestore.MergeAll)
	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return domain.User{}, nil
}

func (repo userRepository) Delete(id string) (bool, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc(id).Delete(ctx)
	if err != nil {
		return false, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return true, nil
}
