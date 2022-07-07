package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			true,
		},
		{
			"Test 2",
			args{
				"aabcc",
				"dbbca",
				"aadbbbaccc",
			},
			false,
		},
		{
			"Test 3",
			args{
				"",
				"",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isInterleave(tt.args.s1, tt.args.s2, tt.args.s3), "isInterleave(%v, %v, %v)", tt.args.s1, tt.args.s2, tt.args.s3)
		})
	}
}
