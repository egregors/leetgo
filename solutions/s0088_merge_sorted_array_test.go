package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)

	nums1 = []int{1}
	nums2 = []int{}
	m, n = 1, 0
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1}, nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	m, n = 0, 1
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1}, nums1)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	m, n = 1, 1
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2}, nums1)
}
