/*
	https://leetcode.com/problems/count-of-smaller-numbers-after-self/

	Given an integer array nums, return an integer array counts where counts[i] is
		the
	number of smaller elements to the right of nums[i].
*/

//nolint:revive // it'ok
package solutions

type BITTree []int

func (t BITTree) Update(i, x, n int) {
	i++
	for i < n {
		t[i] += x
		i += i & -i
	}
}

func (t BITTree) Query(i int) int {
	res := 0
	for i >= 1 {
		res += t[i]
		i -= i & -i
	}
	return res
}

func countSmaller(nums []int) []int {
	offset := 10_000
	size := 2*10_000 + 2
	t := make(BITTree, size)
	var res []int
	for i := len(nums) - 1; i >= 0; i-- {
		smlCount := t.Query(nums[i] + offset)
		res = append(res, smlCount)
		t.Update(nums[i]+offset, 1, size)
	}
	rev := func(xs []int) {
		for i := 0; i < len(xs)/2; i++ {
			xs[i], xs[len(xs)-i-1] = xs[len(xs)-i-1], xs[i]
		}
	}
	rev(res)
	return res
}
