package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"abc", "pqr"},
			"apbqcr",
		},
		{
			"test-2",
			args{"ab", "pqrs"},
			"apbqrs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeAlternately(tt.args.word1, tt.args.word2), "mergeAlternately(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
