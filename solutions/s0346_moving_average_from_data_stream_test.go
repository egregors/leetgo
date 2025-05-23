package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovingAverage(t *testing.T) {
	ma := NewMovingAverage(3)
	assert.Equal(t, 1.0, ma.Next(1))
	assert.Equal(t, 5.5, ma.Next(10))
	assert.Equal(t, 4.666666666666667, ma.Next(3))
	assert.Equal(t, 6.0, ma.Next(5))
}
