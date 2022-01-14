package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "Example 2",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "Example 3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "Example 4",
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "Example 6",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: "Example 7",
			args: args{
				s: "9223372036854775809",
			},
			want: 2147483647,
		},
		{
			name: "Example 8",
			args: args{
				s: "-9223372036854775809",
			},
			want: -2147483648,
		},
		{
			name: "Example 9",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, myAtoi(tt.args.s), "myAtoi(%v)", tt.args.s)
		})
	}
}
