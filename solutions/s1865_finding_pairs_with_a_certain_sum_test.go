package solutions

import (
	"testing"
)

func TestNewFindSumPairs(t *testing.T) {
	pairs := NewFindSumPairs([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	if pairs.Count(7) != 8 {
		t.Errorf("Expected 8, got %d", pairs.Count(7))
	}
	pairs.Add(3, 2)
	if pairs.Count(8) != 2 {
		t.Errorf("Expected 2, got %d", pairs.Count(8))
	}
	if pairs.Count(4) != 1 {
		t.Errorf("Expected 2, got %d", pairs.Count(8))
	}
	pairs.Add(0, 1)
	pairs.Add(1, 1)
	if pairs.Count(7) != 11 {
		t.Errorf("Expected 11, got %d", pairs.Count(7))
	}
}
