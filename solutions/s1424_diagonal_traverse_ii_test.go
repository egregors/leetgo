package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			[]int{1, 4, 2, 7, 5, 3, 8, 6, 9},
		},
		{
			"case 2",
			args{[][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}},
			[]int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findDiagonalOrder(tt.args.nums), "findDiagonalOrder(%v)", tt.args.nums)
		})
	}
}
