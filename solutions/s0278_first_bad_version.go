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
	mid := (min + max) >> 1
	for min <= max {
		if pred(mid) {
			return mid
		} else {
			return bs(mid+1, max, pred)
		}
	}
	return -1
}

func firstBadVersion(n int, isBadVersion func(int) bool) int {
	// step I: find first bad version
	// [. . . . . . . . . . . . . . x x x x x x ]
	// 						^
	// [ . . . . x x x x x x ]
	//				 ^
	// 					max will be first isBad
	// step II: find first good version
	// [ . . . . x x x ]
	// 		   ^
	// 			  min will be first isGood
	// step III: just find first isBad
	// [ . x x x ]
	//     ^
	// 		i is answer
	if n < 10 {
		for i := 1; i <= n; i++ {
			if isBadVersion(i) {
				return i - 1
			}
		}
	}
	bad := bs(1, n, isBadVersion)
	for i := bad; i >= 1; i-- {
		if !isBadVersion(i) {
			return i + 1
		}
	}
	return -1
}
