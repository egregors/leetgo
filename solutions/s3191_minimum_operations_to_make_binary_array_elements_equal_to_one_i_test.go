package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations3191(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{0, 1, 1, 1, 0, 0}},
			3,
		},
		{
			"case 2",
			args{[]int{0, 1, 1, 1}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minOperations3191(tt.args.nums), "minOperations3191(%v)", tt.args.nums)
		})
	}
}
