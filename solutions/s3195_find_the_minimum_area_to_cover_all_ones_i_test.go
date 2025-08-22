package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				grid: [][]int{
					{0, 1, 0},
					{1, 0, 1},
				},
			},
			want: 6,
		},
		{
			name: "Test2",
			args: args{
				grid: [][]int{
					{1, 0},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumArea(tt.args.grid), "minimumArea(%v)", tt.args.grid)
		})
	}
}
