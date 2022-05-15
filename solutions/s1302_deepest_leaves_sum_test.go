package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deepestLeavesSum(t *testing.T) {
	assert.Equal(t, 15, deepestLeavesSum(NewTreeNode("[1,2,3,4,5,null,6,7,null,null,null,null,8]")))
	assert.Equal(t, 19, deepestLeavesSum(NewTreeNode("[6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]")))
}
