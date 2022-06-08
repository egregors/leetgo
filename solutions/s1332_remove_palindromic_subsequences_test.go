package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removePalindromeSub(t *testing.T) {
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
				s: "ababa",
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				s: "abb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removePalindromeSub(tt.args.s), "removePalindromeSub(%v)", tt.args.s)
		})
	}
}
