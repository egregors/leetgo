package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion(t *testing.T) {
	assert.Equal(t, 4, firstBadVersion(5, func(i int) bool {
		return i >= 4
	}))

	assert.Equal(t, 1, firstBadVersion(1, func(i int) bool {
		return i >= 1
	}))

	assert.Equal(t, 1, firstBadVersion(3, func(i int) bool {
		return i >= 1
	}))

	assert.Equal(t, 2, firstBadVersion(3, func(i int) bool {
		return i >= 2
	}))
}
