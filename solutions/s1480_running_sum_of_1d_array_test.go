package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 0",
			args{nums: []int{1, 2, 3, 4}},
			[]int{1, 3, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, runningSum(tt.args.nums), "runningSum(%v)", tt.args.nums)
		})
	}
}
