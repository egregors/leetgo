package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{-1, -8, 0, 5, -9},
			},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSatisfaction(tt.args.satisfaction), "maxSatisfaction(%v)", tt.args.satisfaction)
		})
	}
}
