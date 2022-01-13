package dockerComposeService_test

import (
	mock_ports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/dockerComposeService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	dockerComposeRepository *mock_ports.MockDockerComposeRepository
}

func TestDockerComposeService_Get(t *testing.T) {

	//Mocks//


	id := "1001-1001-1001-1001"

	sampleWantedDockerCompose := domain.DockerCompose{Id: id,DockerComposeDatas: []byte("{id: '1001-1001-1001-1001', value: 'comme ça'") }
	sampleResultDockerCompose := domain.DockerCompose{Id: id,DockerComposeDatas: []byte("{id: '1001-1001-1001-1001', value: 'comme ça'") }


	//Tests//

	type args struct {
		id string
	}

	type want struct {
		result domain.DockerCompose
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should get game successfully",
			args: args{id: "1001-1001-1001-1001"},
			want: want{result: sampleWantedDockerCompose},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.NotFound,nil, "DockerCompose not found in database", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.DockerCompose{}, errors.New(apperrors.NotFound,nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.Internal,nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.DockerCompose{}, errors.New(apperrors.Internal,nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mock_ports.NewMockDockerComposeRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository)

		result, err := service.Get(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}