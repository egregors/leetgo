/*
	https://leetcode.com/problems/find-k-closest-elements/description

		Given a sorted integer array arr, two integers k and x, return the k closest
			integers to x in the array.
	The result should also be sorted in ascending order.

	An integer a is closer to x than an integer b if:

		|a - x| < |b - x|, or
		|a - x| == |b - x| and a < b
*/

package solutions

func findClosestElements(arr []int, k int, x int) []int {
	l, mid, r := 0, 0, len(arr)-k
	for l < r {
		mid = (l + r) / 2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	res := make([]int, 0, k)
	for i := l; i < l+k; i++ {
		res = append(res, arr[i])
	}

	return res
}
