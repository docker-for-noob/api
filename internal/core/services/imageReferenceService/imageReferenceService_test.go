package imageReferenceService_test

import (
	mock_ports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/imageReferenceService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	imageReferenceRepository *mock_ports.MockImageReferenceRepository
}

func TestGetImageReferenceRequest(t *testing.T) {

	//Mocks//
	sampleImageReference := domain.ImageReference{Name: "go:latest",Port : []string{"8080"},  Workdir: []string{"path/to/file"}}

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
			name: "Should get reference successfully",
			args: args{image: "go:latest"},
			want: want{result: sampleImageReference},
			mocks: func(m mockers) {
				m.imageReferenceRepository.EXPECT().Read("go:latest").Return(sampleImageReference, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{image: "go:dernier"},
			want: want{result: domain.ImageReference{},
				err: errors.New(
					apperrors.NotFound,
					nil,
					"Not found",
					"",
				)},
			mocks: func(m mockers) {
				m.imageReferenceRepository.EXPECT().Read("go:dernier").Return(domain.ImageReference{}, nil)
			},
		},
		{
			name: "Should return a internal error",
			args: args{image: "go:dernier"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while searching the reference", "")},
			mocks: func(m mockers) {
				m.imageReferenceRepository.EXPECT().Read("go:dernier").Return(domain.ImageReference{}, errors.New(apperrors.Internal, nil, "", ""))
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
