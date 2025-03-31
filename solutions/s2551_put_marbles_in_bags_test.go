package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_putMarbles(t *testing.T) {
	type args struct {
		weights []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"case 1",
			args{[]int{1, 3, 5, 1}, 2},
			4,
		},
		{
			"case 2",
			args{[]int{1, 3}, 2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, putMarbles(tt.args.weights, tt.args.k), "putMarbles(%v, %v)", tt.args.weights, tt.args.k)
		})
	}
}
