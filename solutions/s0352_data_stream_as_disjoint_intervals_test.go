package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	sr := NewSummaryRanges()
	sr.AddNum(1)
	assert.Equal(t, [][]int{{1, 1}}, sr.GetIntervals())
	sr.AddNum(3)
	assert.Equal(t, [][]int{{1, 1}, {3, 3}}, sr.GetIntervals())
	sr.AddNum(7)
	assert.Equal(t, [][]int{{1, 1}, {3, 3}, {7, 7}}, sr.GetIntervals())
	sr.AddNum(2)
	assert.Equal(t, [][]int{{1, 3}, {7, 7}}, sr.GetIntervals())
	sr.AddNum(6)
	assert.Equal(t, [][]int{{1, 3}, {6, 7}}, sr.GetIntervals())
}
