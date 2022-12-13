package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodingData(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "godevops",
			args: args{
				data: "test-data",
			},
		},
	}
	// TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EncodingData(tt.args.data)
			assert.Equal(t, err, nil)
		})
	}
}
