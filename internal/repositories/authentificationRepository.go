package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	ports "github.com/docker-generator/api/internal/core/ports/user"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"github.com/mitchellh/mapstructure"
)

type authentificationRepository struct {
	BCryptRepository ports.BCryptRepository
}

func NewAuthentificationRepository(BCryptRepository ports.BCryptRepository) *authentificationRepository {
	return &authentificationRepository{
		BCryptRepository: BCryptRepository,
	}
}

func (m authentificationRepository) Login(credentials domain.Credentials) (domain.User, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	snapshot, err := client.Collection("Users").Where("Email", "==", credentials.Email).Snapshots(ctx).Next()

	if snapshot.Size == 0 {
		return domain.User{}, errors.New(apperrors.InvalidInput, err, "User not found", "")
	}

	resp, _ := snapshot.Documents.GetAll()
	userData := resp[0].Data()
	user := domain.User{}

	err = mapstructure.Decode(userData, &user)

	isPasswordValid := m.BCryptRepository.CheckPasswordHash(credentials.Password, user.Password)

	if !isPasswordValid {
		return domain.User{}, errors.New(apperrors.InvalidInput, err, "Invalid password", "")
	}

	_, err = client.Collection("Users").Doc(user.ID).Set(ctx, map[string]interface{}{
		"LastLogin": firestore.ServerTimestamp,
	}, firestore.MergeAll)

	if err != nil {
		return domain.User{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return user, nil
}

func (m authentificationRepository) Logout(id string) error {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc(id).Set(ctx, map[string]interface{}{
		"LastLogout": firestore.ServerTimestamp,
	}, firestore.MergeAll)

	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return nil
}
