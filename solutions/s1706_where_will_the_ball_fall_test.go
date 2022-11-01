package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findBall(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				grid: [][]int{
					{1, 1, 1, -1, -1},
					{1, 1, 1, -1, -1},
					{-1, -1, -1, 1, 1},
					{1, 1, 1, 1, -1},
					{-1, -1, -1, -1, -1},
				},
			},
			want: []int{1, -1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBall(tt.args.grid), "findBall(%v)", tt.args.grid)
		})
	}
}
