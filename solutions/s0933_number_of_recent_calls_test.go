package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRecentCounter(t *testing.T) {
	rc := NewRecentCounter()
	assert.Equal(t, 1, rc.Ping(1))
	assert.Equal(t, 2, rc.Ping(100))
	assert.Equal(t, 3, rc.Ping(3001))
	assert.Equal(t, 3, rc.Ping(3002))
}
