package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countHomogenous(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example 1",
			args{"abbcccaa"},
			13,
		},
		{
			"example 2",
			args{"xy"},
			2,
		},
		{
			"example 3",
			args{"zzzzz"},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countHomogenous(tt.args.s), "countHomogenous(%v)", tt.args.s)
		})
	}
}
