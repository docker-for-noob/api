package formater_test

import (
	"github.com/docker-generator/api/pkg/formater"
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

		result := formater.NormalizeEmail(tt.args)

		assert.Equal(t, tt.want.result, result)
	}
}
