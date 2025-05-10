package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test Case 1",
			args: args{
				nums1: []int{3, 2, 0, 1, 0},
				nums2: []int{6, 5, 0},
			},
			want: 12,
		},
		{
			name: "Test Case 2",
			args: args{
				nums1: []int{2, 0, 2, 0},
				nums2: []int{1, 4},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minSum(tt.args.nums1, tt.args.nums2), "minSum(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
