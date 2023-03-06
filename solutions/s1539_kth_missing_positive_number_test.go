package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				[]int{2, 3, 4, 7, 11},
				5,
			},
			9,
		},
		{
			"Example 2",
			args{
				[]int{1, 2, 3, 4},
				2,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findKthPositive(tt.args.arr, tt.args.k), "findKthPositive(%v, %v)", tt.args.arr, tt.args.k)
		})
	}
}
