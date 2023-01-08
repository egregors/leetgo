package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
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
					{1, 1},
					{2, 2},
					{3, 3},
				},
			},
			3,
		},
		{
			"case 2",
			args{
				[][]int{
					{1, 1},
					{3, 2},
					{5, 3},
					{4, 1},
					{2, 3},
					{1, 4},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxPoints(tt.args.points), "maxPoints(%v)", tt.args.points)
		})
	}
}
