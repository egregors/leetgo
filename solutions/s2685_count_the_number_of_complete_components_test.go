package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countCompleteComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				6,
				[][]int{
					{0, 1},
					{0, 2},
					{1, 2},
					{3, 4},
				},
			},
			3,
		},
		{
			"case 2",
			args{
				6,
				[][]int{
					{0, 1},
					{0, 2},
					{1, 2},
					{3, 4},
					{3, 5},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countCompleteComponents(tt.args.n, tt.args.edges), "countCompleteComponents(%v, %v)", tt.args.n, tt.args.edges)
		})
	}
}
