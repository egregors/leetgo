package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{5, 8, 6},
				3,
			},
			5,
		},
		{
			"test 2",
			args{
				[]int{2, 5},
				11,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumCandies(tt.args.candies, tt.args.k), "maximumCandies(%v, %v)", tt.args.candies, tt.args.k)
		})
	}
}
