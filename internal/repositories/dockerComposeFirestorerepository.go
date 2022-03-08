package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"github.com/mitchellh/mapstructure"
)

type dockerComposeFirestore struct {}

func NewDockerComposeFirestore() *dockerComposeFirestore{
	return &dockerComposeFirestore{}
}

func (repo *dockerComposeFirestore) ReadAll(firstItemRank int) ([]domain.DockerCompose, error ){
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	dockerComposeSearch := client.Collection("Users").Doc("1001-1001-1001-1001").Collection("Files").OrderBy("createdAt", firestore.Desc).Limit(25).Offset(firstItemRank).Documents(ctx)

	dockerComposeItems, err := dockerComposeSearch.GetAll()
	if err != nil {
		return []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	var dockerComposeResults []domain.DockerCompose

	for _, dockerComposeItem := range dockerComposeItems {
		dockerComposeResult := domain.DockerCompose{}
		dockerComposeData := dockerComposeItem.Data()

		err = mapstructure.Decode(dockerComposeData, &dockerComposeResult)
		if err != nil {
			return []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
		}
		dockerComposeResults = append(dockerComposeResults, dockerComposeResult)
	}

	return dockerComposeResults, nil
}

func (repo *dockerComposeFirestore) Read(id string) (domain.DockerCompose, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	dockerComposeSearch, err := client.Collection("Users").Doc("1001-1001-1001-1001").Collection("Files").Doc(id).Get(ctx)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	dockerComposeData := dockerComposeSearch.Data()

	dockerComposeResult := domain.DockerCompose{}

	err = mapstructure.Decode(dockerComposeData, &dockerComposeResult)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	return dockerComposeResult, nil
}
func (repo *dockerComposeFirestore) Create(dockerCompose domain.DockerCompose) (domain.DockerCompose, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc("1001-1001-1001-1001").Collection("Files").Doc(dockerCompose.Id).Set(ctx, map[string]interface{}{
		"id": dockerCompose.Id,
		"name": dockerCompose.Name,
		"DockerComposeDatas": dockerCompose.DockerComposeDatas,
		"createdAt": firestore.ServerTimestamp,
	})
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return domain.DockerCompose{}, nil
}
func (repo *dockerComposeFirestore) Update(dockerCompose domain.DockerCompose) (domain.DockerCompose, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc("1001-1001-1001-1001").Collection("Files").Doc(dockerCompose.Id).Set(ctx, map[string]interface{}{
		"id": dockerCompose.Id,
		"name": dockerCompose.Name,
		"DockerComposeDatas": dockerCompose.DockerComposeDatas,
	}, firestore.MergeAll)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return domain.DockerCompose{}, nil
}

func (repo *dockerComposeFirestore) Delete(id string) (bool, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("Users").Doc("1001-1001-1001-1001").Collection("Files").Doc(id).Delete(ctx)
	if err != nil {
		return false, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return true, nil
}