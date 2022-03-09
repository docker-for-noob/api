package authentificationService_test

import (
	mockports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/authentification/authentificationService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	authentificationRepository *mockports.MockAuthentificationRepository
}

func TestAuthentificationService_Login(t *testing.T) {

	//Mocks//

	sampleCredentials := domain.Credentials{
		Password: "azerty123",
		Email:    "test@test.com",
	}

	id := "1001-1001-1001-1001"

	sampleWantedUser := domain.User{
		ID:       id,
		Email:    sampleCredentials.Email,
		Password: sampleCredentials.Password,
	}

	sampleResultUser := domain.User{
		ID:       id,
		Email:    sampleCredentials.Email,
		Password: sampleCredentials.Password,
	}

	//Tests//

	type want struct {
		result domain.User
		err    error
	}

	tests := []struct {
		name  string
		args  domain.Credentials
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should log in successfully and return a user",
			args: sampleCredentials,
			want: want{result: sampleWantedUser},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Login(sampleCredentials).Return(sampleResultUser, nil)
			},
		},
		{
			name: "Should return an internal error",
			args: sampleCredentials,
			want: want{
				err: errors.New(
					apperrors.Internal,
					nil,
					"An internal error occurred",
					"",
				)},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Login(sampleCredentials).Return(domain.User{}, errors.New(apperrors.Internal, nil, "An internal error occurred", ""))
			},
		},
		{
			name: "Should return a not found error",
			args: sampleCredentials,
			want: want{
				err: errors.New(
					apperrors.NotFound,
					nil, "User not found in database",
					"",
				)},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Login(sampleCredentials).Return(domain.User{}, errors.New(apperrors.NotFound, nil, "User not found in database", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			authentificationRepository: mockports.NewMockAuthentificationRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := authentificationService.New(m.authentificationRepository)

		result, err := service.Login(tt.args)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestAuthentificationService_Logout(t *testing.T) {
	//Mocks//

	id := "1001-1001-1001-1001"

	//Tests//

	type want struct {
		err error
	}
	type args struct {
		id string
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should log out successfully and return an err nil",
			args: args{id: id},
			want: want{err: nil},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Logout(id).Return(nil)
			},
		},
		{
			name: "Should return an internal error",
			args: args{id: id},
			want: want{
				err: errors.New(
					apperrors.Internal,
					nil,
					"An internal error occurred",
					"",
				)},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Logout(id).Return(errors.New(apperrors.Internal, nil, "An internal error occurred", ""))
			},
		},
		{
			name: "Should return a not found error",
			args: args{id: id},
			want: want{
				err: errors.New(
					apperrors.NotFound,
					nil, "User not found in database",
					"",
				)},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Logout(id).Return(errors.New(apperrors.NotFound, nil, "User not found in database", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			authentificationRepository: mockports.NewMockAuthentificationRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := authentificationService.New(m.authentificationRepository)

		err := service.Logout(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		} else {
			assert.Equal(t, tt.want.err, err)
		}
	}
}
