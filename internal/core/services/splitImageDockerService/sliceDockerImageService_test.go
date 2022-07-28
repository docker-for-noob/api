package splitImageDockerService

import (
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/matiasvarela/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitImageRequest(t *testing.T) {

	//Mocks//
	imageDetailForPhpLatest := domain.ImageNameDetail{Name: "php:latest", Language: "php", Version: "latest", Tags: []string{}}
	imageDetailForPhpHuitApache := domain.ImageNameDetail{Name: "php:8.0-apache", Language: "php", Version: "8.0", Tags: []string{"apache"}}
	imageDetailForPhpHuitRcApacheBuster := domain.ImageNameDetail{Name: "php:8.0-rc-apache-buster", Language: "php", Version: "8.0", Tags: []string{"rc", "apache", "buster"}}
	//Tests//

	type args struct {
		imageName string
	}

	type want struct {
		result domain.ImageNameDetail
		err    error
	}

	tests := []struct {
		name  string
		args  args
		want  want
	}{
		{
			name: "Should Split the imahe php:latest",
			args: args{imageName: "php:latest"},
			want: want{result: imageDetailForPhpLatest, err: nil},
		},
		{
			name: "Should Split the imahe php:8.0-apache",
			args: args{imageName: "php:8.0-apache"},
			want: want{result: imageDetailForPhpHuitApache, err: nil},
		},
		{
			name: "Should Split the imahe php:8.0-apache",
			args: args{imageName: "php:8.0-rc-apache-buster"},
			want: want{result: imageDetailForPhpHuitRcApacheBuster, err: nil},
		},
	}

	// Test Runner //

	for _, tt := range tests {
		tt := tt

		result, err := SplitDockerImageName(tt.args.imageName)

		// Verify
		if tt.want.err != nil {
			assert.Equal(t, errors.Code(tt.want.err), errors.Code(err))
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}