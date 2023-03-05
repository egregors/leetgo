package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test 1",
			args{
				[]int{1, 3, 5, 2, 7, 5},
				1,
				5,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubarrays(tt.args.nums, tt.args.minK, tt.args.maxK), "countSubarrays(%v, %v, %v)", tt.args.nums, tt.args.minK, tt.args.maxK)
		})
	}
}
