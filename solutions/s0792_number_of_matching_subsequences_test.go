package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 0",
			args{"abcde", []string{"a", "bb", "acd", "ace"}},
			3,
		},
		{
			"test 1",
			args{"dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numMatchingSubseq(tt.args.s, tt.args.words), "numMatchingSubseq(%v, %v)", tt.args.s, tt.args.words)
		})
	}
}
