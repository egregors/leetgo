package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_diagonalSort(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test 0",
			args{
				[][]int{
					{3, 3, 1, 1},
					{2, 2, 1, 2},
					{1, 1, 1, 2},
				},
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 2},
				{1, 2, 3, 3},
			},
		},
		{
			"Test 1",
			args{mat: [][]int{
				{11, 25, 66, 1, 69, 7},
				{23, 55, 17, 45, 15, 52},
				{75, 31, 36, 44, 58, 8},
				{22, 27, 33, 25, 68, 4},
				{84, 28, 14, 11, 5, 50},
			}},
			[][]int{
				{5, 17, 4, 1, 52, 7},
				{11, 11, 25, 45, 8, 69},
				{14, 23, 25, 44, 58, 15},
				{22, 27, 31, 36, 50, 66},
				{84, 28, 75, 33, 55, 68},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, diagonalSort(tt.args.mat), "diagonalSort(%v)", tt.args.mat)
		})
	}
}
