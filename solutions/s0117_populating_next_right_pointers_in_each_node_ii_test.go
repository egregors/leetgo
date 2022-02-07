package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect117(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 7,
			},
		},
	}
	connect117(root)
	assert.Nil(t, root.Next)
	assert.Equal(t, root.Right, root.Left.Next)
	assert.Nil(t, root.Right.Next)
	assert.Equal(t, root.Left.Right, root.Left.Left.Next)
	assert.Equal(t, root.Right.Right, root.Left.Right.Next)
}
