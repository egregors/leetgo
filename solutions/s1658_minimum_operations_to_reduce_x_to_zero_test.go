package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations1658(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{nums: []int{1, 1, 4, 2, 3}, x: 5},
			2,
		},
		{
			"Test 2",
			args{nums: []int{5, 6, 7, 8, 9}, x: 4},
			-1,
		},
		{
			"Test 3",
			args{nums: []int{3, 2, 20, 1, 1, 3}, x: 10},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minOperations1658(tt.args.nums, tt.args.x), "minOperations1658(%v, %v)", tt.args.nums, tt.args.x)
		})
	}
}
