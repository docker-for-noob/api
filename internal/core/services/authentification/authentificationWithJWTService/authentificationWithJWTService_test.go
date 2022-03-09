package authentificationWithJWTService_test

import (
	mockports "github.com/docker-generator/api/Mocks"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/internal/core/services/authentification/authentificationWithJWTService"
	"github.com/golang/mock/gomock"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockers struct {
	authentificationRepository *mockports.MockAuthentificationRepository
	JWTRepository              *mockports.MockJWTRepository
}

func TestAuthentificationWithJwtService_Login(t *testing.T) {

	//Mocks//

	sampleCredentials := domain.Credentials{
		Password: "azerty123",
		Email:    "test@test.com",
	}

	id := "1001-1001-1001-1001"

	sampleResultUser := domain.User{
		ID:       id,
		Email:    sampleCredentials.Email,
		Password: sampleCredentials.Password,
	}

	sampleResultToken := domain.JwtToken{
		Data: "token",
	}

	//Tests//

	type want struct {
		result domain.JwtToken
		err    error
	}

	tests := []struct {
		name  string
		args  domain.Credentials
		want  want
		mocks func(m mockers)
	}{
		{
			name: "Should log in successfully and return a jwt token",
			args: sampleCredentials,
			want: want{result: sampleResultToken, err: nil},
			mocks: func(m mockers) {
				m.authentificationRepository.EXPECT().Login(sampleCredentials).Return(sampleResultUser, nil)
				m.JWTRepository.EXPECT().CreateJWTTokenString(sampleResultUser).Return(sampleResultToken, nil)
			},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		m := mockers{
			authentificationRepository: mockports.NewMockAuthentificationRepository(gomock.NewController(t)),
			JWTRepository:              mockports.NewMockJWTRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := authentificationWithJWTService.New(m.authentificationRepository, m.JWTRepository)

		result, err := service.Login(tt.args)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}
