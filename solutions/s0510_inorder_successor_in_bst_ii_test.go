package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderSuccessor(t *testing.T) {
	var n1, n2, n3 = new(Node), new(Node), new(Node)
	n1.Val = 1
	n1.Parent = n2

	n2.Val = 2
	n2.Left = n1
	n2.Right = n3

	n3.Val = 3
	n3.Parent = n2

	assert.Equal(t, n2, inorderSuccessor(n1))
}
