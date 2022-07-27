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
	dockerHubRepositoryMock *mock_ports.MockDockerHubRepository
	imageDockerServiceMock *mock_ports.MockImageDockerService
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
			dockerHubRepositoryMock: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			imageDockerServiceMock: mock_ports.NewMockImageDockerService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageReferenceService.New(m.imageReferenceRepository, m.dockerHubRepositoryMock, m.imageDockerServiceMock)

		result, err := service.Get(tt.args.image)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestFindAllTagReferenceForALanguage(t *testing.T) {

	//Mocks//
	imageDockerServiceMockResult := domain.DockerImageResult{
		Name:"go",
		Tags: []string {"dernier", "avantDernier"},

	}
	//Tests//

	type args struct {
		languageName string
	}

	type want struct {
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should return a internal error",
			args: args{languageName: "go"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while searching the reference", "")},
			mocks: func(m mockers) {
				m.imageDockerServiceMock.EXPECT().Get("go", "").Return(imageDockerServiceMockResult, nil)
				m.dockerHubRepositoryMock.EXPECT().HandleMultipleGetTagReference("go",imageDockerServiceMockResult.Tags).Return(errors.New(apperrors.Internal, nil, "", ""))
			},
		},
		{
			name: "Should return a internal error",
			args: args{languageName: "go"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while searching the tags", "")},
			mocks: func(m mockers) {
				m.imageDockerServiceMock.EXPECT().Get("go", "").Return(domain.DockerImageResult{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			imageReferenceRepository: mock_ports.NewMockImageReferenceRepository(gomock.NewController(t)),
			dockerHubRepositoryMock: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			imageDockerServiceMock: mock_ports.NewMockImageDockerService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageReferenceService.New(m.imageReferenceRepository, m.dockerHubRepositoryMock, m.imageDockerServiceMock)

		err := service.FindAllTagReferenceForALanguage(tt.args.languageName)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

	}
}

func TestAddAllTagReference(t *testing.T) {

	//Mocks//
	imageDockerServiceMockResult := domain.DockerImageResult{
		Name:"go",
		Tags: []string {"dernier", "avantDernier"},

	}
	//Tests//

	type args struct {
		allLanguage []string
	}

	type want struct {
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should return a internal error",
			args: args{allLanguage: []string {"go"}},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while searching ALL the reference", "")},
			mocks: func(m mockers) {
				m.imageDockerServiceMock.EXPECT().Get("go", "").Return(domain.DockerImageResult{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
		{
			name: "Should return a internal error",
			args: args{allLanguage: []string {"go"}},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while adding the reference in dadabase", "")},
			mocks: func(m mockers) {
				m.imageDockerServiceMock.EXPECT().Get("go", "").Return(imageDockerServiceMockResult, nil)
				m.dockerHubRepositoryMock.EXPECT().HandleMultipleGetTagReference("go",imageDockerServiceMockResult.Tags).Return(nil)
				m.imageReferenceRepository.EXPECT().AddAllTagReferenceFromApi().Return(errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //
	for _, tt := range tests {
		tt := tt

		m := mockers{
			imageReferenceRepository: mock_ports.NewMockImageReferenceRepository(gomock.NewController(t)),
			dockerHubRepositoryMock: mock_ports.NewMockDockerHubRepository(gomock.NewController(t)),
			imageDockerServiceMock: mock_ports.NewMockImageDockerService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := imageReferenceService.New(m.imageReferenceRepository, m.dockerHubRepositoryMock, m.imageDockerServiceMock)

		err := service.AddAllTagReference(tt.args.allLanguage)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

	}
}