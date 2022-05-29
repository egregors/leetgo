package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			16,
		},
		{
			"Test 1",
			args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			4,
		},
		{
			"Test 2",
			args{words: []string{"a", "aa", "aaa", "aaaa"}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxProduct(tt.args.words), "maxProduct(%v)", tt.args.words)
		})
	}
}
