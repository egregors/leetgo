package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLongestSubsequence(t *testing.T) {
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
			"test1",
			args{
				[]string{"e", "a", "b"},
				[]int{0, 0, 1},
			},
			[]string{"e", "b"},
		},
		{
			"test2",
			args{
				[]string{"a", "b", "c", "d"},
				[]int{1, 0, 1, 1},
			},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getLongestSubsequence(tt.args.words, tt.args.groups), "getLongestSubsequence(%v, %v)", tt.args.words, tt.args.groups)
		})
	}
}
