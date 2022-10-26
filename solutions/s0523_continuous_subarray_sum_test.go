package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkSubarraySum(tt.args.nums, tt.args.k), "checkSubarraySum(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
