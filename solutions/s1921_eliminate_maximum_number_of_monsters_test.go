package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{1, 3, 4},
				[]int{1, 1, 1},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, eliminateMaximum(tt.args.dist, tt.args.speed), "eliminateMaximum(%v, %v)", tt.args.dist, tt.args.speed)
		})
	}
}
