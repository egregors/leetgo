package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSneakyNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: nums = [0,1,1,0]
		//
		//Output: [0,1]
		{
			name: "test case 1",
			args: args{nums: []int{0, 1, 1, 0}},
			want: []int{0, 1},
		},
		//Input: nums = [0,3,2,1,3,2]
		//
		//Output: [2,3]
		{
			name: "test case 2",
			args: args{nums: []int{0, 3, 2, 1, 3, 2}},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderInts(tt.want, getSneakyNumbers(tt.args.nums)), "getSneakyNumbers(%v)", tt.args.nums)
		})
	}
}
