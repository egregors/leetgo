package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				points: [][]int{
					{10, 16},
					{2, 8},
					{1, 6},
					{7, 12},
				},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				points: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
					{7, 8},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMinArrowShots(tt.args.points), "findMinArrowShots(%v)", tt.args.points)
		})
	}
}
