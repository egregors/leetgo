package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_increasingBST(t *testing.T) {
	root := NewTreeNode("[5,3,6,2,4,null,8,1,null,null,null,7,9]")
	assert.Equal(
		t,
		"[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]",
		increasingBST(root).String(),
	)
}
