package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubarrays2962(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test_case_1",
			args: args{
				nums: []int{1, 3, 2, 3, 3},
				k:    2,
			},
			want: 6,
		},
		{
			name: "test_case_2",
			args: args{
				nums: []int{1, 4, 2, 1},
				k:    3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubarrays2962(tt.args.nums, tt.args.k), "countSubarrays2962(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
