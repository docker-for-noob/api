package imageDockerService_test

import (
	mock_ports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/imageDockerService"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	dockerHubRepository *mock_ports.MockDockerHubRepository
	redisRepository     *mock_ports.MockRedisRepository
}

func TestGetImageRequest(t *testing.T) {

	//Mocks//
	sampleWantedDockerImage := domain.DockerImageResult{Name: "php", Tags: []string{"buster", "zt-sbuster"}}

	//Tests//

	type args struct {
		image string
		tag   string
	}

	type want struct {
		result domain.DockerImageResult
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should get image successfully Redis",
			args: args{image: "node", tag: "latest"},
			want: want{result: sampleWantedDockerImage},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().ImageExist("node", "latest").Return(true)
				m.redisRepository.EXPECT().Read("node", "latest").Return(sampleWantedDockerImage, nil)
			},
		},
		{
			name: "Should get image successfully DockerHub",
			args: args{image: "node", tag: "latest"},
			want: want{result: sampleWantedDockerImage},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().ImageExist("node", "latest").Return(false)
				m.dockerHubRepository.EXPECT().Read("node", "latest").Return(sampleWantedDockerImage, nil)
			},
		},
		{
			name: "Should return a NotFound image error",
			args: args{image: "noda", tag: "latest"},
			want: want{result: domain.DockerImageResult{
				Name: "",
				Tags: nil,
			}},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().ImageExist("noda", "latest").Return(false)
				m.dockerHubRepository.EXPECT().Read("noda", "latest").Return(domain.DockerImageResult{}, nil)
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerHubRepository: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			redisRepository:     mock_ports.NewMockRedisRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageDockerService.New(m.dockerHubRepository, m.redisRepository)

		result, err := service.Get(tt.args.image, tt.args.tag)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestGetAllVersionsFromImage(t *testing.T) {

	//Mocks//
	sampleWantedDockerImage := domain.DockerImageResult{Name: "mysql", Tags: []string{"latest", "8.0.30-oracle", "8.0-oracle"}}
	sampleWantedDockerImageWithDoublonVersion := domain.DockerImageResult{Name: "mysql", Tags: []string{"latest", "8.0.30-oracle", "8.0-oracle", "8.0-debian"}}


	//Tests//

	type args struct {
		image string
	}

	type want struct {
		result domain.DockerImageVersions
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should return all version for an image",
			args: args{image: "mysql"},
			want: want{result: domain.DockerImageVersions{Name: "mysql", Versions: []string{"latest", "8.0.30", "8.0"}}},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().ImageExist("mysql", "").Return(true)
				m.redisRepository.EXPECT().Read("mysql", "").Return(sampleWantedDockerImage, nil)
				m.redisRepository.EXPECT().Add("tags_mysql:latest",  []domain.ImageNameDetail{{Name: "mysql:latest", Language: "mysql", Version: "latest", Tags: []string{}}}).AnyTimes()
				m.redisRepository.EXPECT().Add("tags_mysql:8.0.30", []domain.ImageNameDetail{{Name: "mysql:8.0.30-oracle", Language: "mysql", Version: "8.0.30", Tags: []string{"oracle"}}}).AnyTimes()
				m.redisRepository.EXPECT().Add("tags_mysql:8.0", []domain.ImageNameDetail{{Name: "mysql:8.0-oracle", Language: "mysql", Version: "8.0", Tags: []string{"oracle"}}}).AnyTimes()

			},
		},
		{
			name: "Should not return twice the same version for an image",
			args: args{image: "mysql"},
			want: want{result: domain.DockerImageVersions{Name: "mysql", Versions: []string{"latest", "8.0.30", "8.0"}}},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().ImageExist("mysql", "").Return(true)
				m.redisRepository.EXPECT().Read("mysql", "").Return(sampleWantedDockerImageWithDoublonVersion, nil)
				m.redisRepository.EXPECT().Add("tags_mysql:latest",  []domain.ImageNameDetail{{Name: "mysql:latest", Language: "mysql", Version: "latest", Tags: []string{}}}).AnyTimes()
				m.redisRepository.EXPECT().Add("tags_mysql:8.0.30", []domain.ImageNameDetail{{Name: "mysql:8.0.30-oracle", Language: "mysql", Version: "8.0.30", Tags: []string{"oracle"}}}).AnyTimes()
				m.redisRepository.EXPECT().Add("tags_mysql:8.0", []domain.ImageNameDetail{{Name: "mysql:8.0-oracle", Language: "mysql", Version: "8.0", Tags: []string{"oracle"}},{Name: "mysql:8.0-debian", Language: "mysql", Version: "8.0", Tags: []string{"debian"}}}).AnyTimes()
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerHubRepository: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			redisRepository:     mock_ports.NewMockRedisRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageDockerService.New(m.dockerHubRepository, m.redisRepository)

		result, err := service.GetAllVersionsFromImage(tt.args.image)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestGetAllTagsFromImageVersion(t *testing.T) {

	//Mocks//
	//sampleWantedDockerImage := domain.DockerImageResult{Name: "mysql", Tags: []string{"latest", "8.0.30-oracle", "8.0-oracle"}}
	//sampleWantedDockerImageWithDoublonVersion := domain.DockerImageResult{Name: "mysql", Tags: []string{"latest", "8.0.30-oracle", "8.0-oracle", "8.0-debian"}}


	//Tests//

	type args struct {
		languageName string
		version string
	}

	type want struct {
		result []domain.ImageNameDetail
		err error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should return all Tags from cache",
			args: args{languageName: "mysql", version: "5.7.32"},
			want: want{result: []domain.ImageNameDetail{{Name: "mysql:5.7.32", Language: "mysql", Version: "5.7.32", Tags: []string{} }}},
			mocks: func(m mockers) {
				m.redisRepository.EXPECT().FindOne("tags_mysql:5.7.32").Return(`[{"Name": "mysql:5.7.32", "Language": "mysql", "Version": "5.7.32", "Tags": []}]`)
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerHubRepository: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			redisRepository:     mock_ports.NewMockRedisRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageDockerService.New(m.dockerHubRepository, m.redisRepository)

		result, err := service.GetAllTagsFromImageVersion(tt.args.languageName, tt.args.version)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}