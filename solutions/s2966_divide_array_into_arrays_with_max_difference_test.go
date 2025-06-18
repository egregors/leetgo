package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divideArray2966(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test 1",
			args{[]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2},
			[][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			"test 2",
			args{[]int{2, 4, 2, 2, 5, 2}, 2},
			[][]int{},
		},
		{
			"test 3",
			args{[]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14},
			[][]int{{2, 2, 2}, {4, 5, 5}, {5, 5, 7}, {7, 8, 8}, {9, 9, 10}, {11, 12, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, divideArray2966(tt.args.nums, tt.args.k), "divideArray2966(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
