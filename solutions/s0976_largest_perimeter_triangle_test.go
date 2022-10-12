package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{2, 1, 2}},
			5,
		},
		{
			"test-2",
			args{[]int{1, 2, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestPerimeter(tt.args.nums), "largestPerimeter(%v)", tt.args.nums)
		})
	}
}
