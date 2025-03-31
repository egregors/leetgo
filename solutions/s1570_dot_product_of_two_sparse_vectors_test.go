package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseVector_dotProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		vec  SparseVector
		want int
	}{
		{
			"case 1",
			[]int{1, 0, 0, 2, 3},
			NewSparseVector([]int{0, 3, 0, 4, 0}),
			8,
		},
		{
			"case 2",
			[]int{0, 1, 0, 0, 0, 0, 0},
			NewSparseVector([]int{1, 0, 0, 0, 3, 0, 4}),
			0,
		},
		{
			"case 3",
			[]int{0, 1, 0, 0, 2, 0, 0},
			NewSparseVector([]int{1, 0, 0, 0, 3, 0, 4}),
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := NewSparseVector(tt.nums)
			assert.Equalf(t, tt.want, sv.dotProduct(tt.vec), "dotProduct(%v)", tt.vec)
		})
	}
}
