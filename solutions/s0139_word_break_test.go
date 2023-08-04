package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{"leetcode", []string{"leet", "code"}},
			true,
		},
		{
			"case 2",
			args{"applepenapple", []string{"apple", "pen"}},
			true,
		},
		{
			"case 3",
			args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wordBreak(tt.args.s, tt.args.wordDict), "wordBreak(%v, %v)", tt.args.s, tt.args.wordDict)
		})
	}
}
