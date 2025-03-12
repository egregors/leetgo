package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumCount(t *testing.T) {
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
			args{[]int{-2, -1, -1, 1, 2, 3}},
			3,
		},
		{
			"case 2",
			args{[]int{-3, -2, -1, 0, 0, 1, 2}},
			3,
		},
		{
			"case 3",
			args{[]int{5, 20, 66, 1314}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumCount(tt.args.nums), "maximumCount(%v)", tt.args.nums)
		})
	}
}
