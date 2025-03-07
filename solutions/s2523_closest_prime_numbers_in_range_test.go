package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{2, 10},
			[]int{2, 3},
		},
		{
			"Test 2",
			args{10, 19},
			[]int{11, 13},
		},
		{
			"Test 3",
			args{4, 6},
			[]int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, closestPrimes(tt.args.left, tt.args.right), "closestPrimes(%v, %v)", tt.args.left, tt.args.right)
		})
	}
}
