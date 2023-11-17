package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPairSum(t *testing.T) {
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
			args{[]int{3, 5, 2, 3}},
			7,
		},
		{
			"case 2",
			args{[]int{3, 5, 4, 2, 4, 6}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minPairSum(tt.args.nums), "minPairSum(%v)", tt.args.nums)
		})
	}
}
