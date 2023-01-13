package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPath(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{
				[]int{-1, 0, 0, 1, 1, 2},
				"abacbe",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPath(tt.args.parent, tt.args.s), "longestPath(%v, %v)", tt.args.parent, tt.args.s)
		})
	}
}
