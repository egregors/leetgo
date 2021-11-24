/*
	https://leetcode.com/problems/first-bad-version/

	You are a product manager and currently leading a team to develop a new product.
	Unfortunately, the latest version of your product fails the quality check.
	Since each version is developed based on the previous version, all the versions after a bad version are also bad.

	Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one,
	which causes all the following ones to be bad.

	You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a
	function to find the first bad version. You should minimize the number of calls to the API.
*/

package solutions

func bs(min, max int, pred func(int) bool) int {
	var mid int
	for min <= max {
		mid = (min + max) >> 1
		if pred(mid) {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	if pred(mid) {
		return mid
	}
	return mid + 1
}

func firstBadVersion(n int, isBadVersion func(int) bool) int {
	return bs(1, n, isBadVersion)
}
