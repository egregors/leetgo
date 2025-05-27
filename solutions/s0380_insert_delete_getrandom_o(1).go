/*
	https://leetcode.com/problems/insert-delete-getrandom-o1/description/

	Implement the RandomizedSet class:

    RandomizedSet() Initializes the RandomizedSet object.
		bool insert(int val) Inserts an item val into the set if not present. Returns
			true if the item was not present,
	false otherwise.
		bool remove(int val) Removes an item val from the set if present. Returns true
			if the item was present, false
	otherwise.
		int getRandom() Returns a random element from the current set of elements
			(it's guaranteed that at least one
	element exists when this method is called). Each element must have the same
		probability of being returned.

	You must implement the functions of the class such that each function works in
		average O(1) time complexity.
*/

package solutions

import "math/rand"

type RandomizedSet struct {
	xs []int       // vals
	m  map[int]int // val : idx in xs
}

// NewRandomizedSet must call Constructor to pass LeetCode test
func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{xs: make([]int, 0, 16), m: make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; !ok {
		s.xs = append(s.xs, val)
		s.m[val] = len(s.xs) - 1

		return true
	}

	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.m[val]; ok {
		// [1 2 3 4]
		//.   ^
		// 2 ~ 4
		// [1 4 3 2] -> pop last -> [1 4 3]
		if len(s.xs) != 1 {
			l, r := s.m[val], len(s.xs)-1
			s.m[s.xs[r]] = l
			s.xs[l], s.xs[r] = s.xs[r], s.xs[l]
		}

		delete(s.m, val)
		s.xs = s.xs[:len(s.xs)-1]

		return true
	}

	return false
}

func (s *RandomizedSet) GetRandom() int {
	if len(s.xs) == 1 {
		return s.xs[0]
	}

	return s.xs[rand.Intn(len(s.xs))] //nolint:gosec
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
