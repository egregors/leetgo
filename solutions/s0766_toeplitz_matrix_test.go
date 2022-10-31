package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{
				[][]int{
					{1, 2, 3, 4},
					{5, 1, 2, 3},
					{9, 5, 1, 2},
				},
			},
			true,
		},
		{
			"case 2",
			args{
				[][]int{
					{1, 2},
					{2, 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isToeplitzMatrix(tt.args.matrix), "isToeplitzMatrix(%v)", tt.args.matrix)
		})
	}
}
