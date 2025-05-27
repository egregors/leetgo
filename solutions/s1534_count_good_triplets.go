/*
	https://leetcode.com/problems/count-good-triplets/description/

	Given an array of integers arr, and three integers a, b and c. You need to find
		the number of
	good triplets.

	A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are
		true:

		0 <= i < j < k < arr.length
		|arr[i] - arr[j]| <= a
		|arr[j] - arr[k]| <= b
		|arr[i] - arr[k]| <= c

	Where |x| denotes the absolute value of x.

	Return the number of good triplets.
*/

package solutions

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	n := len(arr)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if Abs(arr[i]-arr[j]) <= a &&
					Abs(arr[j]-arr[k]) <= b &&
					Abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}

	return count
}
