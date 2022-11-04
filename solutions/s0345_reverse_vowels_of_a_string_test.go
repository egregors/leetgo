package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"hello"},
			"holle",
		},
		{
			"2",
			args{"leetcode"},
			"leotcede",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseVowels(tt.args.s), "reverseVowels(%v)", tt.args.s)
		})
	}
}
