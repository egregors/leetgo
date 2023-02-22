package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				5,
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shipWithinDays(tt.args.weights, tt.args.days), "shipWithinDays(%v, %v)", tt.args.weights, tt.args.days)
		})
	}
}
