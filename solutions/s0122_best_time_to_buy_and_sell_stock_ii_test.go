package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit0122(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxProfit0122(tt.args.prices), "maxProfit0122(%v)", tt.args.prices)
		})
	}
}
