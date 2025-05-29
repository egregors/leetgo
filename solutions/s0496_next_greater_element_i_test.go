package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
		// Output: [-1,3,-1]
		{
			name: "example1",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		// Input: nums1 = [2,4], nums2 = [1,2,3,4]
		// Output: [3,-1]
		{
			name: "example2",
			args: args{
				nums1: []int{2, 4},
				nums2: []int{1, 2, 3, 4},
			},
			want: []int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreaterElement(tt.args.nums1, tt.args.nums2), "nextGreaterElement(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
