package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
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
			name: "Test 1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},

		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, subarraySum(tt.args.nums, tt.args.k), "subarraySum(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
