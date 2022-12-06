package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	assert.Equal(t, []int{1, 3, 5, 2, 4}, oddEvenList(NewListNode([]int{1, 2, 3, 4, 5})).ToSlice())
	assert.Equal(t, []int{2, 6, 1, 4, 3, 7, 5}, oddEvenList(NewListNode([]int{2, 3, 6, 7, 1, 5, 4})).ToSlice())
}
