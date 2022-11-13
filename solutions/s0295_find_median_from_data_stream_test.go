package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMedianFinder2(t *testing.T) {
	mf := NewMedianFinder()
	mf.AddNum(41)
	assert.Equal(t, 41.0, mf.FindMedian())
	mf.AddNum(35)
	assert.Equal(t, 38.0, mf.FindMedian())
	mf.AddNum(62)
	assert.Equal(t, 41.0, mf.FindMedian())
	mf.AddNum(5)
	assert.Equal(t, 38.0, mf.FindMedian())
	mf.AddNum(97)
	assert.Equal(t, 41.0, mf.FindMedian())
	mf.AddNum(108)
	assert.Equal(t, 51.5, mf.FindMedian())
}
func TestNewMedianFinder1(t *testing.T) {
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian())
	mf.AddNum(3)
	assert.Equal(t, 2.0, mf.FindMedian())
}
