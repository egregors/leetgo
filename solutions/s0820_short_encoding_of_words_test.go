package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumLengthEncoding(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"time", "me", "bell"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumLengthEncoding(tt.args.words), "minimumLengthEncoding(%v)", tt.args.words)
		})
	}
}
