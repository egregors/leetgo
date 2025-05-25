package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{[]int{1}},
			true,
		},
		{
			"test case 2",
			args{[]int{2, 1, 4}},
			true,
		},
		{
			"test case 3",
			args{[]int{4, 3, 1, 6}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isArraySpecial(tt.args.nums), "isArraySpecial(%v)", tt.args.nums)
		})
	}
}
