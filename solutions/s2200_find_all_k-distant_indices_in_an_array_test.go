package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: nums = [3,4,9,1,3,9,5], key = 9, k = 1
		//Output: [1,2,3,4,5,6]
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 4, 9, 1, 3, 9, 5},
				key:  9,
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		//Input: nums = [2,2,2,2,2], key = 2, k = 2
		//Output: [0,1,2,3,4]
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				key:  2,
				k:    2,
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k), "findKDistantIndices(%v, %v, %v)", tt.args.nums, tt.args.key, tt.args.k)
		})
	}
}
