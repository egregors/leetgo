package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arraySign(t *testing.T) {
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
			args{[]int{-1, -2, -3, -4, 3, 2, 1}},
			1,
		},
		{
			"case 2",
			args{[]int{1, 5, 0, 2, -3}},
			0,
		},
		{
			"case 3",
			args{[]int{-1, 1, -1, 1, -1}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, arraySign(tt.args.nums), "arraySign(%v)", tt.args.nums)
		})
	}
}
