package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{s: "(()"},
			2,
		},
		{
			"Test 2",
			args{s: ")()())"},
			4,
		},
		{
			"Test 3",
			args{s: ""},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestValidParentheses(tt.args.s), "longestValidParentheses(%v)", tt.args.s)
		})
	}
}
