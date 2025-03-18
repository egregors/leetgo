package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestNiceSubarray(t *testing.T) {
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
			args{[]int{1, 3, 8, 48, 10}},
			3,
		},
		{
			"case 2",
			args{[]int{3, 1, 5, 11, 13}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestNiceSubarray(tt.args.nums), "longestNiceSubarray(%v)", tt.args.nums)
		})
	}
}
