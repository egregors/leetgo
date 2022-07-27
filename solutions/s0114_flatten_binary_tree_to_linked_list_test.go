package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	root := NewTreeNode("[1,2,5,3,4,null,6]")
	flatten(root)
	assert.Equal(t, "[1,null,2,null,3,null,4,null,5,null,6]", root.String())
}
