package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedArrayToBST(t *testing.T) {
	res := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	assert.Equal(t, "[0,-3,9,-10,null,5,null]", res.String())
}
