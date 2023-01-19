package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraysDivByK(t *testing.T) {
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
			"test 1",
			args{
				nums: []int{4, 5, 0, -2, -3, 1},
				k:    5,
			},
			7,
		},
		{
			"test 2",
			args{
				nums: []int{5},
				k:    9,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, subarraysDivByK(tt.args.nums, tt.args.k), "subarraysDivByK(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
