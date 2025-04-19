package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "abca",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validPalindrome(tt.args.s), "validPalindrome(%v)", tt.args.s)
		})
	}
}
