/*
	Given an integer array nums, return true if there exists a triple of indices (i, j, k)
	such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
*/

package solutions

func increasingTriplet(nums []int) bool {
	maxInt := 1<<31 - 1
	fst, snd := maxInt, maxInt

	for _, n := range nums {
		if n <= fst {
			fst = n
		} else if n <= snd {
			snd = n
		} else {
			return true
		}
	}

	return false
}
