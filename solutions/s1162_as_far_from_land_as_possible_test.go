package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDistance(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				[][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxDistance(tt.args.grid), "maxDistance(%v)", tt.args.grid)
		})
	}
}
