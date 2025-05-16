package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test 1",
			args{
				words:  []string{"a", "b", "c", "d"},
				groups: []int{1, 2, 3, 4},
			},
			[]string{"a", "b", "c", "d"},
		},
		{
			"Test 2",
			args{
				[]string{"bab", "dab", "cab"},
				[]int{1, 2, 2},
			},
			[]string{"bab", "dab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getWordsInLongestSubsequence(tt.args.words, tt.args.groups), "getWordsInLongestSubsequence(%v, %v)", tt.args.words, tt.args.groups)
		})
	}
}
