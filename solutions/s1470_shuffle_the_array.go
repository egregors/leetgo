/*
	https://leetcode.com/problems/shuffle-the-array/

	Given the array nums consisting of 2n elements in the form
	[x1,x2,...,xn,y1,y2,...,yn].

	Return the array in the form [x1,y1,x2,y2,...,xn,yn].
*/

package solutions

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		res = append(res, nums[i], nums[i+n])
	}
	return res
}
