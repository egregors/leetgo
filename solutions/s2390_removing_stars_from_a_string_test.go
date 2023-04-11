package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty string",
			args{""},
			"",
		},
		{
			"string without stars",
			args{"abc"},
			"abc",
		},
		{
			"string with stars",
			args{"a*b*c"},
			"c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeStars(tt.args.s), "removeStars(%v)", tt.args.s)
		})
	}
}
