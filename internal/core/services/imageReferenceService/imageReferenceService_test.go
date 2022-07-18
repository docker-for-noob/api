package imageReferenceService_test

import (
	mock_ports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/imageReferenceService"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	imageReferenceRepository *mock_ports.MockImageReferenceRepository
}

func TestGetImageRequest(t *testing.T) {

	//Mocks//
	sampleImageReference := domain.ImageReference{Name: "php", Workdir: []string{"buster", "zt-sbuster"}}

	//Tests//

	type args struct {
		image string
		tag   string
	}

	type want struct {
		result domain.ImageReference
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
			args: args{image: "node", tag: "latest"},
			want: want{result: sampleImageReference},
			mocks: func(m mockers) {
				m.imageReferenceRepository.EXPECT().Read("node").Return(sampleImageReference, nil)
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			imageReferenceRepository: mock_ports.NewMockImageReferenceRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageReferenceService.New(m.imageReferenceRepository)

		result, err := service.Get(tt.args.image)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}
