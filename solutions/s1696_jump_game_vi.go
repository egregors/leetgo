/*
	https://leetcode.com/problems/jump-game-vi/

	You are given a 0-indexed integer array nums and an integer k.

	You are initially standing at index 0. In one move, you can jump at most k
		steps
	forward without going outside the boundaries of the array. That is, you can
		jump
	from index i to any index in the range [i + 1, min(n - 1, i + k)] inclusive.

	You want to reach the last index of the array (index n - 1). Your score is the
		sum of
	all nums[j] for each index j you visited in the array.

	Return the maximum score you can get.
*/

package solutions

func maxResult(nums []int, k int) int {
	n := len(nums)
	score := make([]int, n)
	score[0] = nums[0]
	dq := Deque[int]{}
	dq.PushLast(0)

	for i := 1; i < n; i++ {
		for !dq.IsEmpty() && dq.PeekFirst() < i-k {
			dq.PopFirst()
		}
		score[i] = score[dq.PeekFirst()] + nums[i]
		for !dq.IsEmpty() && score[i] >= score[dq.PeekLast()] {
			dq.PopLast()
		}
		dq.PushLast(i)
	}
	return score[n-1]
}
