/*
	You are given an integer array nums. You want to maximize the number of points you get
	by performing the following operation any number of times:

		Pick any nums[i] and delete it to earn nums[i] points. Afterwards, you must delete
	every element equal to nums[i] - 1 and every element equal to nums[i] + 1.

	Return the maximum number of points you can earn by applying the above operation some number
	of times.
*/

package solutions

func deleteAndEarn(nums []int) int {
	ps := make(map[int]int)
	maxNumber := 0

	for _, num := range nums {
		ps[num] += num
		maxNumber = Maximum(num, maxNumber)
	}

	maxPs := make([]int, maxNumber+1)
	maxPs[1] = ps[1]

	for num := 2; num < len(maxPs); num++ {
		gain := ps[num]
		maxPs[num] = Maximum(maxPs[num-1], maxPs[num-2]+gain)
	}
	return maxPs[maxNumber]
}
