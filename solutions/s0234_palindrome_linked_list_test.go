package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome234(t *testing.T) {
	assert.True(t, isPalindrome234(NewListNode([]int{1, 2, 2, 1})))
	assert.False(t, isPalindrome234(NewListNode([]int{1, 2, 2, 3})))
}
