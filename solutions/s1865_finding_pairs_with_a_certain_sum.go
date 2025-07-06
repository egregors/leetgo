/*
	https://leetcode.com/problems/finding-pairs-with-a-certain-sum

	You are given two integer arrays nums1 and nums2. You are tasked to implement a
		data structure that supports queries of two types:

		Add a positive integer to an element of a given index in the array nums2.
		Count the number of pairs (i, j) such that nums1[i] + nums2[j] equals a given
			value (0 <= i < nums1.length and 0 <= j < nums2.length).

	Implement the FindSumPairs class:

		FindSumPairs(int[] nums1, int[] nums2) Initializes the FindSumPairs object
			with two integer arrays nums1 and nums2.
		void add(int index, int val) Adds val to nums2[index], i.e., apply
			nums2[index] += val.
		int count(int tot) Returns the number of pairs (i, j) such that nums1[i] +
			nums2[j] == tot.

*/

package solutions

type FindSumPairs struct {
	nums1, nums2 []int
	m2           map[int]int
}

func NewFindSumPairs(nums1 []int, nums2 []int) FindSumPairs {
	m2 := make(map[int]int)
	for _, n := range nums2 {
		m2[n]++
	}
	return FindSumPairs{
		m2:    m2,
		nums1: nums1,
		nums2: nums2,
	}
}

func (f2 *FindSumPairs) Add(index int, val int) {
	x := f2.nums2[index]
	f2.m2[x]--
	if f2.m2[x] == 0 {
		delete(f2.m2, x)
	}
	f2.nums2[index] += val
	f2.m2[f2.nums2[index]]++
}

func (f2 *FindSumPairs) Count(tot int) int {
	cnt := 0
	for _, n := range f2.nums1 {
		if v2, ok := f2.m2[tot-n]; ok {
			cnt += v2
		}
	}
	return cnt
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
