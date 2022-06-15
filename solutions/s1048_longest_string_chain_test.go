package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{words: []string{"a", "b", "ba", "bca", "bda", "bdca"}},
			4,
		},
		{
			"Test 2",
			args{words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestStrChain(tt.args.words), "longestStrChain(%v)", tt.args.words)
		})
	}
}
