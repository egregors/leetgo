package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	kthLargest := NewKthLargest(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kthLargest.Add(3))
	assert.Equal(t, 5, kthLargest.Add(5))
	assert.Equal(t, 5, kthLargest.Add(10))
	assert.Equal(t, 8, kthLargest.Add(9))
	assert.Equal(t, 8, kthLargest.Add(4))
}
