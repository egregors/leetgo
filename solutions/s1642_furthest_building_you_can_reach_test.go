package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_furthestBuilding(t *testing.T) {
	type args struct {
		heights []int
		bricks  int
		ladders int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				heights: []int{4, 2, 7, 6, 9, 14, 12},
				bricks:  5,
				ladders: 1,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				heights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19},
				bricks:  10,
				ladders: 2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, furthestBuilding(tt.args.heights, tt.args.bricks, tt.args.ladders), "furthestBuilding(%v, %v, %v)", tt.args.heights, tt.args.bricks, tt.args.ladders)
		})
	}
}
