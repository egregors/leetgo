package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test 1",
			args{[]int{4, 6, 7, 7}},
			[][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
		},
		{
			"test 2",
			args{[]int{4, 4, 3, 2, 1}},
			[][]int{{4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findSubsequences(tt.args.nums), "findSubsequences(%v)", tt.args.nums)
		})
	}
}
