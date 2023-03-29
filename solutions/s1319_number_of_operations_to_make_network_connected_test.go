package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeConnected(t *testing.T) {
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
			"test 1",
			args{
				4,
				[][]int{
					{0, 1},
					{0, 2},
					{1, 2},
				}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeConnected(tt.args.n, tt.args.connections), "makeConnected(%v, %v)", tt.args.n, tt.args.connections)
		})
	}
}
