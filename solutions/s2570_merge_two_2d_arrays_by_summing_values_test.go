package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeArrays(t *testing.T) {
	type args struct {
		nums1 [][]int
		nums2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				nums1: [][]int{{1, 2}, {2, 3}, {4, 5}},
				nums2: [][]int{{1, 4}, {2, 3}, {4, 5}},
			},
			want: [][]int{{1, 6}, {2, 6}, {4, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeArrays(tt.args.nums1, tt.args.nums2), "mergeArrays(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
