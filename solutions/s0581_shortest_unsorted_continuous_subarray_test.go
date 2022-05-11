package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 0",
			args{nums: []int{2, 6, 4, 8, 10, 9, 15}},
			5,
		},
		{
			"test 1",
			args{nums: []int{1, 2, 3, 4}},
			0,
		},
		{
			"test 2",
			args{nums: []int{1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findUnsortedSubarray(tt.args.nums), "findUnsortedSubarray(%v)", tt.args.nums)
		})
	}
}
