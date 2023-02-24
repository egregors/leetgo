/*
	https://leetcode.com/problems/minimize-deviation-in-array

	You are given an array nums of n positive integers.

	You can perform two types of operations on any element of the array any number
	of times:

		If the element is even, divide it by 2.
			For example, if the array is [1,2,3,4], then you can do this operation
	on the last element, and the array will be [1,2,3,2].
		If the element is odd, multiply it by 2.
			For example, if the array is [1,2,3,4], then you can do this operation
	on the first element, and the array will be [2,2,3,4].

	The deviation of the array is the maximum difference between any two elements
	in the array.

	Return the minimum deviation the array can have after performing some number
	of operations.
*/

package solutions

import (
	"container/heap"
	"math"
)

func minimumDeviation(nums []int) int {
	n, minNum, deviation := len(nums), math.MaxInt32, math.MaxInt32
	h := make(IntMaxHeap, n)
	for i := range nums {
		if nums[i]%2 == 0 {
			h[i] = nums[i]
		} else {
			h[i] = nums[i] * 2
		}
		minNum = min(minNum, h[i])
	}
	heap.Init(&h)
	for {
		maxNum := heap.Pop(&h).(int)
		deviation = min(deviation, maxNum-minNum)
		if maxNum%2 == 1 {
			break
		}
		maxNum >>= 1
		heap.Push(&h, maxNum)
		minNum = Minimum(minNum, maxNum)
	}
	return deviation
}
