package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth1614(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{"(1+(2*3)+((8)/4))+1"},
			3,
		}, {
			"case 2",
			args{"(1)+((2))+(((3)))"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxDepth1614(tt.args.s), "maxDepth1614(%v)", tt.args.s)
		})
	}
}
