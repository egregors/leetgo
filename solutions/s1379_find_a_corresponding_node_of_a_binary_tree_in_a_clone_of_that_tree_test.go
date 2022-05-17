package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getTargetCopy(t *testing.T) {
	t1 := NewTreeNode("[7,4,3,null,null,6,19]")
	t1V := FindTreeNode(t1, 3)

	t2 := NewTreeNode("[7,4,3,null,null,6,19]")
	t2V := FindTreeNode(t2, 3)

	assert.Equal(t, t2V, getTargetCopy(t1, t2, t1V))
}
