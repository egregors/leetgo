package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				n: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPowerOfFour(tt.args.n), "isPowerOfFour(%v)", tt.args.n)
		})
	}
}
