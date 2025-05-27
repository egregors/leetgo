/*
	https://leetcode.com/problems/reordered-power-of-2/

	You are given an integer n. We reorder the digits in any order (including the
		original order)
	such that the leading digit is not zero.

	Return true if and only if we can do this so that the resulting number is a
		power of two.
*/

package solutions

func reorderedPowerOf2(n int) bool {
	maxPO2 := 1
	for maxPO2 < 1_000_000_000 {
		maxPO2 *= 2
	}
	ifPO2 := func(n int) bool { return maxPO2%n == 0 }

	var bt func(curr int, xs []int, seen map[int]bool, acc *[]int)
	bt = func(curr int, xs []int, seen map[int]bool, acc *[]int) {
		if seen[curr] {
			return
		}

		if len(xs) == 0 {
			*acc = append(*acc, curr)
			seen[curr] = true
			return
		}

		for j := 0; j < len(xs); j++ {
			x := xs[j]
			if curr == 0 && x == 0 {
				// skip if the leading digit is zero
				continue
			}
			curr *= 10
			curr += x
			cp := make([]int, len(xs))
			copy(cp, xs)
			bt(curr, append(xs[:j], xs[j+1:]...), seen, acc)
			xs = cp
			curr -= x
			curr /= 10
		}
	}
	var nums []int
	bt(0, toSlice(n), make(map[int]bool), &nums)
	for _, num := range nums {
		if ifPO2(num) {
			return true
		}
	}
	return false
}

func toSlice(x int) []int {
	var res []int
	for x/10 > 0 {
		res = append(res, x%10)
		x /= 10
	}
	res = append(res, x)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
