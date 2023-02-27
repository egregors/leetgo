package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node427
	}{
		{
			"example 1",
			args{
				[][]int{
					{0, 1},
					{1, 0},
				},
			},
			&Node427{
				IsLeaf: false,
				Val:    true,
				TopLeft: &Node427{
					IsLeaf: true,
					Val:    false,
				},
				TopRight: &Node427{
					IsLeaf: true,
					Val:    true,
				},
				BottomLeft: &Node427{
					IsLeaf: true,
					Val:    true,
				},
				BottomRight: &Node427{
					IsLeaf: true,
					Val:    false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, construct(tt.args.grid), "construct(%v)", tt.args.grid)
		})
	}
}
