package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				6,
				[][]int{
					{0, 1},
					{1, 3},
					{2, 3},
					{4, 0},
					{4, 5},
				}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minReorder(tt.args.n, tt.args.connections), "minReorder(%v, %v)", tt.args.n, tt.args.connections)
		})
	}
}
