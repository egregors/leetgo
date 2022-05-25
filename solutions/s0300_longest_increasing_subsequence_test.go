package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},
		{
			"Test 1",
			args{nums: []int{0, 1, 0, 3, 2, 3}},
			4,
		},
		{
			"Test 2",
			args{nums: []int{7, 7, 7, 7, 7, 7, 7}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLIS(tt.args.nums), "lengthOfLIS(%v)", tt.args.nums)
		})
	}
}
