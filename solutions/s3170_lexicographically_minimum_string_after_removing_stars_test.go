package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clearStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: s = "aaba*"
		//Output: "aab"
		{
			"t1",
			args{"aaba*"},
			"aab",
		},
		//Input: s = "abc"
		//Output: "abc"
		{
			"t2",
			args{"abc"},
			"abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, clearStars(tt.args.s), "clearStars(%v)", tt.args.s)
		})
	}
}
