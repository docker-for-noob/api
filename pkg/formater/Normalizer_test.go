package formater_test

import (
	formatService "github.com/docker-generator/api/internal/core/services/formatService"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthentificationService_Login(t *testing.T) {
	//Mocks//

	sampleEmail := "   TEST@TEST.COM     "
	sampleResultEmail := "test@test.com"

	//Tests//

	type want struct {
		result string
	}

	tests := []struct {
		name string
		args string
		want want
	}{
		{
			name: "Should return a normalized email",
			args: sampleEmail,
			want: want{result: sampleResultEmail},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		formater := formatService.New()
		result := formater.NormalizeEmail(tt.args)

		assert.Equal(t, tt.want.result, result)
	}
}
