package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partitionArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{[]int{3, 6, 1, 2, 5}, 2},
			2,
		},
		{
			"Test 2",
			args{[]int{1, 2, 3}, 1},
			2,
		},
		{
			"Test 3",
			args{[]int{2, 2, 4, 5}, 0},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, partitionArray(tt.args.nums, tt.args.k), "partitionArray(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
