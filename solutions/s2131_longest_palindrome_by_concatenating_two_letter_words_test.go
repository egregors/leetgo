package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome2131(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				words: []string{"lc", "cl", "gg"},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				words: []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			},
			want: 8,
		},
		{
			name: "test3",
			args: args{
				words: []string{"cc", "ll", "xx"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindrome2131(tt.args.words), "longestPalindrome2131(%v)", tt.args.words)
		})
	}
}
