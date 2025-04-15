package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			name: "test2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderInts(tt.want, intersection(tt.args.nums1, tt.args.nums2)), "intersection(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
