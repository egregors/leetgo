package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nearestExit(t *testing.T) {
	type args struct {
		maze     [][]byte
		entrance []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]byte{
					{'+', '+', '.', '+'},
					{'.', '.', '.', '+'},
					{'+', '+', '+', '.'},
				},
				[]int{1, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nearestExit(tt.args.maze, tt.args.entrance), "nearestExit(%v, %v)", tt.args.maze, tt.args.entrance)
		})
	}
}
