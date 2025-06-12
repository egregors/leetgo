package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxAdjacentDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1, 2, 4}},
			3,
		},
		{
			"2",
			args{[]int{-5, -10, -5}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxAdjacentDistance(tt.args.nums), "maxAdjacentDistance(%v)", tt.args.nums)
		})
	}
}
