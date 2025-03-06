package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	tests := []struct {
		name string
		g    [][]int
		want []int
	}{
		{
			"test 1",
			[][]int{{1, 3}, {2, 2}},
			[]int{2, 4},
		},
		{
			"test 2",
			[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}},
			[]int{9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMissingAndRepeatedValues(tt.g), "findMissingAndRepeatedValues(%v)", tt.g)
		})
	}
}
