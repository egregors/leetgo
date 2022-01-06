package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	want := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	assert.Equal(t, want, combine(4, 2))
}
