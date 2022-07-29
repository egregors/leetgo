package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test0",
			args{words: []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, pattern: "abb"},
			[]string{"mee", "aqq"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findAndReplacePattern(tt.args.words, tt.args.pattern), "findAndReplacePattern(%v, %v)", tt.args.words, tt.args.pattern)
		})
	}
}
