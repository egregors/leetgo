package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkValidCuts(t *testing.T) {
	type args struct {
		n          int
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{
				5,
				[][]int{
					{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkValidCuts(tt.args.n, tt.args.rectangles), "checkValidCuts(%v, %v)", tt.args.n, tt.args.rectangles)
		})
	}
}
