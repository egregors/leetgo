/*
	https://leetcode.com/problems/smallest-number-in-infinite-set/

	You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

	Implement the SmallestInfiniteSet class:

		SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain
			all positive integers.
		int popSmallest() Removes and returns the smallest integer contained in the
			infinite set.
		void addBack(int num) Adds a positive integer num back into the infinite set,
			if it is not already in the infinite set.

*/

//nolint:revive // it's ok
package solutions

type SmallestInfiniteSet struct {
	x int
	m map[int]bool
}

// NewSmallestInfiniteSet should call Constructor to pass LeetCode test
func NewSmallestInfiniteSet() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		x: 1,
		m: make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	n := s.x
	for s.m[n] {
		s.x++
		n++
	}
	s.m[n] = true
	s.x++
	return n
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.x {
		s.x = num
	}
	s.m[num] = false
}
