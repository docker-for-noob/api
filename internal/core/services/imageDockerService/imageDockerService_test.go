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
