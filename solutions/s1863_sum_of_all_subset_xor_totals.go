/*
	https://leetcode.com/problems/sum-of-all-subset-xor-totals/description/

	The XOR total of an array is defined as the bitwise XOR of all its elements, or 0 if the array is empty.

		For example, the XOR total of the array [2,5,6] is 2 XOR 5 XOR 6 = 1.

	Given an array nums, return the sum of all XOR totals for every subset of nums.

	Note: Subsets with the same elements should be counted multiple times.

	An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b.
*/

package solutions

import (
	"math"
)

func subsetXORSum(nums []int) int {
	subs := make([][]int, 0, int(math.Pow(2, float64(len(nums)))))

	var bt func(int, []int)
	bt = func(i int, cur []int) {
		subs = append(subs, cur)
		for i := i; i < len(nums); i++ {
			next := make([]int, len(cur))
			copy(next, cur)
			bt(i+1, append(next, nums[i]))
		}
	}

	bt(0, []int{})

	sum := 0
	for _, sub := range subs {
		subSum := 0
		for _, s := range sub {
			subSum ^= s
		}
		sum += subSum
	}

	return sum
}
