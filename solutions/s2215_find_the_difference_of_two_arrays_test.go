package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{

			"case 0",
			args{
				[]int{1, 2, 3},
				[]int{2, 4, 6},
			},
			[][]int{
				{1, 3},
				{4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(
				t,
				IsEqualAnyOrderIntsSlices(tt.want, findDifference(tt.args.nums1, tt.args.nums2)),
				"findDifference(%v, %v)",
				tt.args.nums1,
				tt.args.nums2,
			)
		})
	}
}
