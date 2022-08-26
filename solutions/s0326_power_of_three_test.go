package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 0",
			args{n: 27},
			true,
		},
		{
			"Test 1",
			args{n: 0},
			false,
		},
		{
			"Test 2",
			args{n: 9},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPowerOfThree(tt.args.n), "isPowerOfThree(%v)", tt.args.n)
		})
	}
}
