package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: nums = [1,2,2,1,1,0]
		//Output: [1,4,2,0,0,0]
		{
			name: "Test 1",
			args: args{nums: []int{1, 2, 2, 1, 1, 0}},
			want: []int{1, 4, 2, 0, 0, 0},
		},
		//Input: nums = [0,1]
		//Output: [1,0]
		{
			name: "Test 2",
			args: args{nums: []int{0, 1}},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, applyOperations(tt.args.nums), "applyOperations(%v)", tt.args.nums)
		})
	}
}
