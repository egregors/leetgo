package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{"abab"},
			true,
		},
		{
			"case 2",
			args{"aba"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, repeatedSubstringPattern(tt.args.s), "repeatedSubstringPattern(%v)", tt.args.s)
		})
	}
}
