package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				n: 218,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPowerOfTwo(tt.args.n), "isPowerOfTwo(%v)", tt.args.n)
		})
	}
}
