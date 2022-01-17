package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Test 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "Test 3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Test 4",
			args: args{
				s: "ab",
			},
			want: "a",
		},
		{
			name: "Test 5",
			args: args{
				s: "abc",
			},
			want: "a",
		},
		{
			name: "Test 6",
			args: args{
				s: "abcd",
			},
			want: "a",
		},
		{
			name: "Test 7",
			args: args{
				s: "abcdc",
			},
			want: "cdc",
		},
		{
			name: "Test 8",
			args: args{
				s: "abcdcb",
			},
			want: "bcdcb",
		},
		{
			name: "Test 9",
			args: args{
				s: "abcdcba",
			},
			want: "abcdcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindrome5(tt.args.s), "longestPalindrome5(%v)", tt.args.s)
		})
	}
}
