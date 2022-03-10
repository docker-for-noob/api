package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/docker-generator/api/internal/core/domain"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/matiasvarela/errors"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

type versionFirestore struct{}

func NewVersionFirestore() *versionFirestore {
	return &versionFirestore{}
}

func (repo *versionFirestore) Create(versionDatas domain.DockerCompose, userId string) error {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	versionId := time.Now().Unix()

	_, err := client.Collection("Users").Doc(userId).Collection("Files").Doc(versionDatas.Id).Collection("Version").Doc(strconv.FormatInt(versionId, 10)).Set(ctx, map[string]interface{}{
		"id":                 versionId,
		"name":               versionDatas.Name,
		"DockerComposeDatas": versionDatas.DockerComposeDatas,
		"createdAt":          firestore.ServerTimestamp,
	})
	if err != nil {
		return errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return nil
}

func (repo *versionFirestore) Read(idDockerCompose string, idVersion string, userId string) (domain.DockerCompose, error) {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	versionSearch, err := client.Collection("Users").Doc(userId).Collection("Files").Doc(idDockerCompose).Collection("Version").Doc(idVersion).Get(ctx)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}
	versionData := versionSearch.Data()

	versionData["id"] = idVersion

	versionResult := domain.DockerCompose{}

	err = mapstructure.Decode(versionData, &versionResult)
	if err != nil {
		return domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	return versionResult, nil
}

func (repo *versionFirestore) ReadAll(idDockerCompose string, userId string) ([]domain.DockerCompose, error) {

	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	allVersionsSearch := client.Collection("Users").Doc(userId).Collection("Files").Doc(idDockerCompose).Collection("Version").OrderBy("createdAt", firestore.Desc).Documents(ctx)

	versionsItems, err := allVersionsSearch.GetAll()
	if err != nil {
		return []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
	}

	var versionResults []domain.DockerCompose

	for _, versionItem := range versionsItems {
		versionResult := domain.DockerCompose{}
		versionData := versionItem.Data()

		idVersion := versionData["id"].(int64)
		versionData["id"] = strconv.FormatInt(idVersion, 10)

		err = mapstructure.Decode(versionData, &versionResult)
		if err != nil {
			return []domain.DockerCompose{}, errors.New(apperrors.Internal, err, "An internal error occurred", "")
		}
		versionResults = append(versionResults, versionResult)
	}

	return versionResults, nil
}
