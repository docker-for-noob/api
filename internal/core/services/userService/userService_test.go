package userService_test

import (
	mockports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/userService"
	apperrors "github.com/docker-generator/api/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	userRepository              *mockports.MockUserRepository
	BCryptRepository            *mockports.MockBCryptRepository
	passwordValidatorRepository *mockports.MockPasswordValidatorRepository
	uidgen                      *mockports.MockUIDGen
}

func TestUserService_Get(t *testing.T) {

	//Mocks//

	id := "1001-1001-1001-1001"

	sampleCredentials := domain.Credentials{
		Password: "azerty123",
		Email:    "test@test.com",
	}

	id = "1001-1001-1001-1001"

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
			name: "Should get a user successfully",
			args: args{id: "1001-1001-1001-1001"},
			want: want{result: sampleWantedUser},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Read("1001-1001-1001-1001").Return(sampleResultUser, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.NotFound, nil, "User not found in database", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.User{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error",
			args: args{id: "1001-1001-1001-1001"},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Read("1001-1001-1001-1001").Return(domain.User{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			userRepository:              mockports.NewMockUserRepository(gomock.NewController(t)),
			BCryptRepository:            mockports.NewMockBCryptRepository(gomock.NewController(t)),
			passwordValidatorRepository: mockports.NewMockPasswordValidatorRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := userService.New(m.userRepository, m.BCryptRepository, m.passwordValidatorRepository, m.uidgen)
		result, err := service.Get(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
		assert.Equal(t, tt.want.result, result)
	}
}

func TestUserService_Patch(t *testing.T) {

	//Mocks//

	id := "1001-1001-1001-1001"

	sampleCredentials := domain.Credentials{
		Password: "azerty123",
		Email:    "test@test.com",
	}

	id = "1001-1001-1001-1001"

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

	sampleInputUser := domain.User{
		ID:       id,
		Email:    sampleCredentials.Email,
		Password: sampleCredentials.Password,
	}
	//Tests//

	type want struct {
		result domain.User
		err    error
	}
	type args struct {
		user domain.User
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should update a User successfully",
			args: args{user: sampleInputUser},
			want: want{result: sampleWantedUser},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Update(sampleInputUser.ID, sampleInputUser).Return(sampleResultUser, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{user: sampleInputUser},
			want: want{err: errors.New(apperrors.NotFound, nil, "User not found in database", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Update(sampleInputUser.ID, sampleInputUser).Return(domain.User{}, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error",
			args: args{user: sampleInputUser},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Update(sampleInputUser.ID, sampleInputUser).Return(domain.User{}, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			userRepository:              mockports.NewMockUserRepository(gomock.NewController(t)),
			BCryptRepository:            mockports.NewMockBCryptRepository(gomock.NewController(t)),
			passwordValidatorRepository: mockports.NewMockPasswordValidatorRepository(gomock.NewController(t)),
			uidgen:                      mockports.NewMockUIDGen(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := userService.New(m.userRepository, m.BCryptRepository, m.passwordValidatorRepository, m.uidgen)
		result, err := service.Patch(tt.args.user.ID, tt.args.user)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
		assert.Equal(t, tt.want.result, result)
	}
}

func TestUserService_Post(t *testing.T) {

	//Mocks//

	sampleCredentials := domain.Credentials{
		Password: "Azerty1234567@",
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
	type args struct {
		user domain.User
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should Create a User successfully",
			args: args{user: sampleWantedUser},
			want: want{result: sampleResultUser},
			mocks: func(m mockers) {
				m.passwordValidatorRepository.EXPECT().VerifyPasswordStrenght(sampleWantedUser.Password).Return(nil)
				m.BCryptRepository.EXPECT().HashPassword(sampleWantedUser.Password).Return(sampleResultUser.Password, nil)
				m.uidgen.EXPECT().NewUuid().Return(id)
				m.userRepository.EXPECT().Create(sampleWantedUser).Return(sampleResultUser, nil)
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			userRepository:              mockports.NewMockUserRepository(gomock.NewController(t)),
			BCryptRepository:            mockports.NewMockBCryptRepository(gomock.NewController(t)),
			passwordValidatorRepository: mockports.NewMockPasswordValidatorRepository(gomock.NewController(t)),
			uidgen:                      mockports.NewMockUIDGen(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := userService.New(m.userRepository, m.BCryptRepository, m.passwordValidatorRepository, m.uidgen)
		result, err := service.Post(tt.args.user)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result.ID, id)
		assert.Equal(t, tt.want.result.Email, result.Email)
		assert.Equal(t, tt.want.result.Password, result.Password)
	}
}

func TestUserService_Delete(t *testing.T) {

	//Mocks//

	id := "1001-1001-1001-1001"

	//Tests//

	type want struct {
		result bool
		err    error
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
			name: "Should delete a User successfully",
			args: args{id: id},
			want: want{result: true},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Delete(id).Return(true, nil)
			},
		},
		{
			name: "Should return a NotFound error",
			args: args{id: id},
			want: want{err: errors.New(apperrors.NotFound, nil, "User not found in database", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Delete(id).Return(false, errors.New(apperrors.NotFound, nil, "", ""))
			},
		},
		{
			name: "Should return a Internal error",
			args: args{id: id},
			want: want{err: errors.New(apperrors.Internal, nil, "An internal error occurred", "")},
			mocks: func(m mockers) {
				m.userRepository.EXPECT().Delete(id).Return(false, errors.New(apperrors.Internal, nil, "", ""))
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			userRepository:              mockports.NewMockUserRepository(gomock.NewController(t)),
			BCryptRepository:            mockports.NewMockBCryptRepository(gomock.NewController(t)),
			passwordValidatorRepository: mockports.NewMockPasswordValidatorRepository(gomock.NewController(t)),
			uidgen:                      mockports.NewMockUIDGen(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := userService.New(m.userRepository, m.BCryptRepository, m.passwordValidatorRepository, m.uidgen)
		result, err := service.Delete(tt.args.id)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}
		assert.Equal(t, tt.want.result, result)
	}
}
