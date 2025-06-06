package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_robotWithString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: s = "zza"
		//Output: "azz"
		{
			"1",
			args{"zza"},
			"azz",
		},
		//Input: s = "bac"
		//Output: "abc"
		{
			"2",
			args{"bac"},
			"abc",
		},
		//Input: s = "bdda"
		//Output: "addb"
		{
			"3",
			args{"bdda"},
			"addb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, robotWithString(tt.args.s), "robotWithString(%v)", tt.args.s)
		})
	}
}
