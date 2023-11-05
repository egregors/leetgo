package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getWinner(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{2, 1, 3, 5, 4, 6, 7},
				2,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getWinner(tt.args.arr, tt.args.k), "getWinner(%v, %v)", tt.args.arr, tt.args.k)
		})
	}
}
