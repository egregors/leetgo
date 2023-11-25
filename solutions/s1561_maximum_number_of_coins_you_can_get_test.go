package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxCoins(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 0",
			args{[]int{2, 4, 1, 2, 7, 8}},
			9,
		},
		{
			"case 1",
			args{[]int{2, 4, 5}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxCoins(tt.args.piles), "maxCoins(%v)", tt.args.piles)
		})
	}
}
