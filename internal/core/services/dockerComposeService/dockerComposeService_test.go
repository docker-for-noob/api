package dockerComposeService_test

import (
	mockports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/dockerComposeService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	dockerComposeRepository *mockports.MockDockerComposeRepository
	versionService *mockports.MockVersionService
}

func TestDockerComposeService_GetAll(t *testing.T) {

	//Mocks//
	firstItem := 0
	id := "1001-1001-1001-1001"
	sampleWantedDockerCompose := []domain.DockerCompose{
		{Id: id, DockerComposeDatas: "{id: '1001-1001-1001-1001', value: 'comme ça'"},
		{Id: id, DockerComposeDatas: "{id: '2002-2002-2002-2002', value: 'comme ça'"},
	}

	sampleResultDockerCompose := []domain.DockerCompose{
		{Id: id, DockerComposeDatas: "{id: '1001-1001-1001-1001', value: 'comme ça'"},
		{Id: id, DockerComposeDatas: "{id: '2002-2002-2002-2002', value: 'comme ça'"},
	}
	//Tests//
	type args struct {
		firstItemRank int
	}

	type want struct {
		lastItemRank int
		result       []domain.DockerCompose
		err          error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should get all dockerCompose successfully",
			args: args{firstItemRank: 0},
			want: want{
				result: sampleWantedDockerCompose,
				lastItemRank: 25,
			},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().ReadAll(firstItem).Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "Should return an internal error",
			args: args{firstItemRank: 0},
			want: want{
				result: []domain.DockerCompose{},
				lastItemRank: 0,
				err: errors.New(apperrors.Internal, nil, "An internal error occurred", ""),
			},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().ReadAll(firstItem).Return([]domain.DockerCompose{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
	}

	// Test Runner //
	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionService: mockports.NewMockVersionService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository, m.versionService)

		lastItemRank, result, err := service.GetAll(tt.args.firstItemRank)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
		assert.Equal(t, tt.want.lastItemRank, lastItemRank)
	}
}

func TestDockerComposeService_Get(t *testing.T) {

	//Mocks//
	id := "1001-1001-1001-1001"
	sampleWantedDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{id: '1001-1001-1001-1001', value: 'comme ça'"}
	sampleResultDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{id: '1001-1001-1001-1001', value: 'comme ça'"}

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
			name: "Should get dockerCompose successfully",
			args: args{id: "1001-1001-1001-1001"},
			want: want{result: sampleWantedDockerCompose},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.NotFound, nil, "DockerCompose not found in database", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.DockerCompose{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //
	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionService: mockports.NewMockVersionService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository, m.versionService)

		result, err := service.Get(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestDockerComposeService_Post(t *testing.T) {

	//Mocks//

	id := "1001-1001-1001-1001"

	sampleWantedDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}
	sampleResultDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}
	sampleInputDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}

	//Tests//

	type args struct {
		dockerCompose domain.DockerCompose
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
			name: "Should Create dockerCompose successfully",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{result: sampleWantedDockerCompose},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Create(sampleInputDockerCompose).Return(sampleResultDockerCompose, nil)
			},
		},
		{
			name: "Should return an Internal error",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Create(sampleInputDockerCompose).Return(domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionService: mockports.NewMockVersionService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository, m.versionService)
		result, err := service.Post(tt.args.dockerCompose)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}

func TestDockerComposeService_Patch(t *testing.T) {

	//Mocks//

	id := "1001-1001-1001-1001"

	sampleWantedDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}
	sampleResultDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}
	sampleInputDockerCompose := domain.DockerCompose{Id: id, DockerComposeDatas: "{value: 'comme ça'"}

	//Tests//

	type args struct {
		dockerCompose domain.DockerCompose
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
			name: "Should Update dockerCompose successfully",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{result: sampleWantedDockerCompose},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Update(sampleInputDockerCompose).Return(sampleResultDockerCompose, nil)
				m.versionService.EXPECT().Add(sampleInputDockerCompose.Id).Return(nil)
			},
		},
		{
			name: "Should return an Internal error",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Update(sampleInputDockerCompose).Return(domain.DockerCompose{}, errors.New(apperrors.Internal, nil, "", ""))
				m.versionService.EXPECT().Add(sampleInputDockerCompose.Id).Return(nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{err: errors.New(apperrors.NotFound, nil, "DockerCompose not found in database", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Update(sampleInputDockerCompose).Return(domain.DockerCompose{}, errors.New(apperrors.NotFound,nil, "", ""))
				m.versionService.EXPECT().Add(sampleInputDockerCompose.Id).Return(nil)
			},
		},
		{
			name: "Should return a NotFound error because Id not found in versionService",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{err: errors.New(apperrors.NotFound, nil, "version Service can not found dockerCompose", "")},
			mocks: func(m mockers) {
				m.versionService.EXPECT().Add(sampleInputDockerCompose.Id).Return(errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error from versionService",
			args: args{dockerCompose: sampleInputDockerCompose},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred in versionService", "")},
			mocks: func(m mockers) {
				m.versionService.EXPECT().Add(sampleInputDockerCompose.Id).Return(errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionService: mockports.NewMockVersionService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository, m.versionService)
		result, err := service.Patch(tt.args.dockerCompose)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}

func TestDockerComposeService_Delete(t *testing.T) {

	//Mocks//

	//Tests//

	type args struct {
		id string
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should Create dockerCompose successfully",
			args: args{id: "1001-1001-1001-1001"},
			want: want{result: true},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Delete("1001-1001-1001-1001").Return(true, nil)
			},
		},
		{
			name: "Should return an Internal error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Delete("1001-1001-1001-1001").Return(false, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.NotFound, nil, "DockerCompose not found in database", "")},
			mocks: func(m mockers) {
				m.dockerComposeRepository.EXPECT().Delete("1001-1001-1001-1001").Return(false, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			dockerComposeRepository: mockports.NewMockDockerComposeRepository(gomock.NewController(t)),
			versionService: mockports.NewMockVersionService(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := dockerComposeService.New(m.dockerComposeRepository, m.versionService)
		result, err := service.Delete(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}
