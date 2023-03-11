package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 01",
			args{
				[]int{3, 6, 7, 11},
				8,
			},
			4,
		},

		{
			"case 02",
			args{
				[]int{30, 11, 23, 4, 20},
				5,
			},
			30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minEatingSpeed(tt.args.piles, tt.args.h), "minEatingSpeed(%v, %v)", tt.args.piles, tt.args.h)
		})
	}
}
