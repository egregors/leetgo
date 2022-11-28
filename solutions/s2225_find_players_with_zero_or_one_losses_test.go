package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test 1",
			args{
				[][]int{
					{1, 3},
					{2, 3},
					{3, 6},
					{5, 6},
					{5, 7},
					{4, 5},
					{4, 8},
					{4, 9},
					{10, 4},
					{10, 9},
				},
			},
			[][]int{
				{1, 2, 10},
				{4, 5, 7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findWinners(tt.args.matches), "findWinners(%v)", tt.args.matches)
		})
	}
}
