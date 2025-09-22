/*
	https://leetcode.com/problems/count-elements-with-maximum-frequency/

	You are given an array nums consisting of positive integers.

	Return the total frequencies of elements in nums such that those elements all
		have the maximum frequency.

	The frequency of an element is the number of occurrences of that element in the
		array.
*/

package solutions

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	mx := 0
	for _, n := range nums {
		m[n]++
		mx = max(mx, m[n])
	}

	s := 0
	for _, v := range m {
		if v == mx {
			s += v
		}
	}

	return s
}
