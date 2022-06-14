package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumUniqueSubarray(t *testing.T) {
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
			args{nums: []int{4, 2, 4, 5, 6}},
			17,
		},
		{
			"Test 1",
			args{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumUniqueSubarray(tt.args.nums), "maximumUniqueSubarray(%v)", tt.args.nums)
		})
	}
}
