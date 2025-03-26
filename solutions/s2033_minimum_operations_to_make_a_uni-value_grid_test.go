package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations2033(t *testing.T) {
	type args struct {
		grid [][]int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[][]int{
					{2, 4},
					{6, 8},
				},
				2,
			},
			4,
		},
		{
			"case 2",
			args{
				[][]int{
					{1, 5},
					{2, 3},
				},
				1,
			},
			5,
		},
		{
			"case 3",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				2,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minOperations2033(tt.args.grid, tt.args.x), "minOperations2033(%v, %v)", tt.args.grid, tt.args.x)
		})
	}
}
