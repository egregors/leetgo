package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion(t *testing.T) {
	assert.Equal(t, 4, firstBadVersion(5, func(i int) bool {
		if i >= 4 {
			return true
		}
		return false
	}))

	assert.Equal(t, 1, firstBadVersion(1, func(i int) bool {
		if i == 1 {
			return true
		}
		return false
	}))
}
