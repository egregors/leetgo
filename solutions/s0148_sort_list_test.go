package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	l := NewListNode([]int{4, 2, 1, 3})
	sorted := sortList(l)
	assert.Equal(t, []int{1, 2, 3, 4}, sorted.ToSlice())
}
