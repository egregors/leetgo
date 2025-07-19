package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations2357(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Input: nums = [1,5,0,3,5]
		//Output: 3
		//Input: nums = [0]
		//Output: 0
		{
			name: "Test1",
			args: args{nums: []int{1, 5, 0, 3, 5}},
			want: 3,
		},
		{
			name: "Test2",
			args: args{nums: []int{0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumOperations2357(tt.args.nums), "minimumOperations2357(%v)", tt.args.nums)
		})
	}
}
