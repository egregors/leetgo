package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"example 1",
			args{
				[]string{"hello", "leetcode"},
				"hlabcdefgijkmnopqrstuvwxyz",
			},
			true,
		},
		{
			"example 2",
			args{
				[]string{"word", "world", "row"},
				"worldabcefghijkmnpqstuvxyz",
			},
			false,
		},
		{
			"example 3",
			args{
				[]string{"apple", "app"},
				"abcdefghijklmnopqrstuvwxyz",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isAlienSorted(tt.args.words, tt.args.order), "isAlienSorted(%v, %v)", tt.args.words, tt.args.order)
		})
	}
}
