package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPaths(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				7,
				[][]int{
					{0, 6, 7},
					{0, 1, 2},
					{1, 2, 3},
					{1, 3, 3},
					{6, 3, 3},
					{3, 5, 1},
					{6, 5, 1},
					{2, 5, 1},
					{0, 4, 5},
					{4, 6, 2},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countPaths(tt.args.n, tt.args.roads), "countPaths(%v, %v)", tt.args.n, tt.args.roads)
		})
	}
}
