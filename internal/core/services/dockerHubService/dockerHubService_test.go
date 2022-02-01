package dockerHubService_test

import (
	mock_ports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/dockerHubService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	dockerHubRepository *mock_ports.MockDockerHubRepository
}

func TestGetImageRequest(t *testing.T) {

	//Mocks//
	sampleWantedDockerHub := domain.DockerHub{DockerHubData: []byte("{name: 'node', namespace: 'library'")}

	//Tests//

	type args struct {
		image string
	}

	type want struct {
		result domain.DockerHub
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should get image successfully",
			args: args{image: "node"},
			want: want{result: sampleWantedDockerHub},
			mocks: func(m mockers) {
				m.dockerHubRepository.EXPECT().Read("node").Return(sampleWantedDockerHub, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{image: "noda"},
			want: want{err: errors.New(apperrors.NotFound, nil, "Image not found in docker hub", "")},
			mocks: func(m mockers) {
				m.dockerHubRepository.EXPECT().Read("noda").Return(domain.DockerHub{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerHubRepository: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerHubService.New(m.dockerHubRepository)

		result, err := service.Get(tt.args.image)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}
