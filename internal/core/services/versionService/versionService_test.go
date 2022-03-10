package versionService_test

import (
	mockports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/versionService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	dockerComposeRepository *mockports.MockDockerComposeRepository
	versionrepository       *mockports.MockVersionRepository
}

func TestVersionService_Add(t *testing.T) {

	id := "1001-1001-1001-1001"
	userId := "1001"

	sampleResultDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}

	type args struct {
		id     string
		userId string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should add a dockerCompose version successfully",
			args: args{id: id, userId: userId},
			want: want{err: nil},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read(id, userId).Return(sampleResultDockerCompose, nil)
				m.versionrepository.EXPECT().Create(sampleResultDockerCompose, userId).Return(nil)
			},
		},
		{
			name: "dockerCompose Read return a not found error",
			args: args{id: id, userId: userId},
			want: want{err: errors.New(apperrors.NotFound, nil, "previous docker-compose version not found", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read(id, userId).Return(domain.DockerCompose{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "dockerCompose Read return a internal error",
			args: args{id: id, userId: userId},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while searching the pervious version", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read(id, userId).Return(domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
		{
			name: "Version create return a internal error",
			args: args{id: id, userId: userId},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occured while creating the version", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read(id, userId).Return(sampleResultDockerCompose, nil)
				m.versionrepository.EXPECT().Create(sampleResultDockerCompose, userId).Return(errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionrepository:       mockports.NewMockVersionRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := versionService.New(m.dockerComposeRepository, m.versionrepository)
		err := service.Add(tt.args.id, tt.args.userId)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		if tt.want.err == nil {
			assert.Equal(t, tt.want.err, err)
		}

	}
}

func TestVersionService_Get(t *testing.T) {

	id := "1001-1001-1001-1001"
	idVersion := "1646256391"
	userId := "1001"

	sampleResultDockerCompose := domain.DockerCompose{Id: idVersion, DockerComposeDatas: "{value: 'comme ça'"}

	type args struct {
		id        string
		idVersion string
		userId    string
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
			name: "Should add a dockerCompose version successfully",
			args: args{id: id, idVersion: idVersion, userId: userId},
			want: want{result: sampleResultDockerCompose, err: nil},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().Read(id, idVersion, userId).Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "repository Read return a not found error",
			args: args{id: id, idVersion: idVersion, userId: userId},
			want: want{result: domain.DockerCompose{}, err: errors.New(apperrors.NotFound, nil, "version not found", "")},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().Read(id, idVersion, userId).Return(domain.DockerCompose{}, nil)
			},
		},
		{
			name: "repository Read return a internal error",
			args: args{id: id, idVersion: idVersion, userId: userId},
			want: want{result: domain.DockerCompose{}, err: errors.New(apperrors.Internal, nil, "an error occured while searching the version", "")},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().Read(id, idVersion, userId).Return(domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "an error occured while searching the version", ""))
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionrepository:       mockports.NewMockVersionRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := versionService.New(m.dockerComposeRepository, m.versionrepository)
		result, err := service.Get(tt.args.id, tt.args.idVersion, tt.args.userId)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}

func TestVersionService_GetAll(t *testing.T) {

	id := "1001-1001-1001-1001"
	userId := "1001"
	sampleResultDockerCompose := []domain.DockerCompose{
		{Id: "1646256391", DockerComposeDatas: "{id: '1646256391', value: 'comme ça'"},
		{Id: "1646254757", DockerComposeDatas: "{id: '1646254757', value: 'comme ça'"},
	}

	type args struct {
		id     string
		userId string
	}

	type want struct {
		result []domain.DockerCompose
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should add a dockerCompose version successfully",
			args: args{id: id, userId: userId},
			want: want{result: sampleResultDockerCompose, err: nil},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().ReadAll(id, userId).Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "service ReadAll should return a not found error",
			args: args{id: id, userId: userId},
			want: want{result: []domain.DockerCompose{}, err: errors.New(apperrors.NotFound, nil, "version not found", "")},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().ReadAll(id, userId).Return([]domain.DockerCompose{}, nil)
			},
		},
		{
			name: "repository Read return a internal error",
			args: args{id: id, userId: userId},
			want: want{result: []domain.DockerCompose{}, err: errors.New(apperrors.Internal, nil, "an error occured while searching the version", "")},
			mocks: func(m mockers) {
				m.versionrepository.EXPECT().ReadAll(id, userId).Return([]domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "an error occured while searching the version", ""))
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionrepository:       mockports.NewMockVersionRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := versionService.New(m.dockerComposeRepository, m.versionrepository)
		result, err := service.GetAll(tt.args.id, tt.args.userId)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}
