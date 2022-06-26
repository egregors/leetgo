/*
	https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

	There are several cards arranged in a row, and each card has an associated number of points.
	The points are given in the integer array cardPoints.

	In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.

	Your score is the sum of the points of the cards you have taken.

	Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
*/

package solutions

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	front, back := make([]int, k+1), make([]int, k+1)

	for i := 0; i < k; i++ {
		front[i+1] = cardPoints[i] + front[i]
		back[i+1] = cardPoints[n-1-i] + back[i]
	}

	maxScore := 0
	for i := 0; i <= k; i++ {
		curr := front[i] + back[k-i]
		maxScore = Maximum(maxScore, curr)
	}

	return maxScore
}
