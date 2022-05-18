package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_criticalConnections(t *testing.T) {
	assert.Equal(t, [][]int{{1, 3}}, criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
	assert.Equal(t, [][]int{{0, 1}}, criticalConnections(2, [][]int{{0, 1}}))
}
