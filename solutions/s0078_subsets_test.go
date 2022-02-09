package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test 1",
			args{nums: []int{1, 2, 3}},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, subsets(tt.args.nums), "subsets(%v)", tt.args.nums)
		})
	}
}
