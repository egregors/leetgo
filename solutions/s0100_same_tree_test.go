package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	assert.True(t, isSameTree(NewTreeNode("[1,2,3]"), NewTreeNode("[1,2,3]")))
	assert.False(t, isSameTree(NewTreeNode("[1,2]"), NewTreeNode("[1,null,2]")))
	assert.False(t, isSameTree(NewTreeNode("[1,2,1]"), NewTreeNode("[1,1,2]")))
}
