package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	assert.Equal(t, []int{1, 2, 2, 4, 3, 5}, partition(NewListNode([]int{1, 4, 3, 2, 5, 2}), 3).ToSlice())
}
