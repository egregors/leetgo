package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"e", "o"}},
			[]string{"facebook", "google", "leetcode"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wordSubsets(tt.args.words1, tt.args.words2), "wordSubsets(%v, %v)", tt.args.words1, tt.args.words2)
		})
	}
}
