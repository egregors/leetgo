package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{0, 1, 0},
					{1, 1, 1},
					{0, 1, 0},
				},
				target: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numSubmatrixSumTarget(tt.args.matrix, tt.args.target), "numSubmatrixSumTarget(%v, %v)", tt.args.matrix, tt.args.target)
		})
	}
}
