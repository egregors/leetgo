package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				word1: "sea",
				word2: "eat",
			},
			want: 2,
		},
		{
			name: "Example 1",
			args: args{
				word1: "leetcode",
				word2: "etco",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDistance(tt.args.word1, tt.args.word2), "minDistance(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
