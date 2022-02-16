/*
	https://leetcode.com/problems/third-maximum-number/

	Given an integer array nums, return the third distinct maximum number in this array.
	If the third maximum does not exist, return the maximum number.
*/

//nolint:revive // it's ok
package solutions

type Num404s map[int]bool

func (n Num404s) Max() int {
	max := -1 << 31
	for k := range n {
		if k > max {
			max = k
		}
	}
	return max
}

func thirdMax(nums []int) int {
	m := make(Num404s)

	// O(n)
	for _, n := range nums {
		m[n] = true
	}

	if len(m) < 3 {
		return m.Max()
	}

	// O(n*3) ~> O(n)
	for i := 0; i < 2; i++ {
		delete(m, m.Max())
	}

	return m.Max()
}
