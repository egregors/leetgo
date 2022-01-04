package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				grid: [][]int{
					{0, 2},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, orangesRotting(tt.args.grid), "orangesRotting(%v)", tt.args.grid)
		})
	}
}
