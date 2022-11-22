package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case 1",
			args{"aab"},
			"aba",
		},
		{
			"case 2",
			args{"aaab"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reorganizeString(tt.args.s), "reorganizeString(%v)", tt.args.s)
		})
	}
}
