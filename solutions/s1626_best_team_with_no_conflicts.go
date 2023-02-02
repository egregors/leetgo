/*
	https://leetcode.com/problems/best-team-with-no-conflicts/

	You are the manager of a basketball team. For the upcoming tournament, you want to choose
	the team with the highest overall score. The score of the team is the sum of scores of
	all the players in the team.

	However, the basketball team is not allowed to have conflicts. A conflict exists if a
	younger player has a strictly higher score than an older player. A conflict does not
	occur between players of the same age.

	Given two lists, scores and ages, where each scores[i] and ages[i] represents the score
	and age of the ith player, respectively, return the highest overall score of all possible
	basketball teams.
*/

package solutions

import "sort"

type pair1626 struct {
	score int
	age   int
}

func bestTeamScore(scores []int, ages []int) int {

	sa := make([]pair1626, 0, len(scores))

	for i := 0; i < len(scores); i++ {
		sa = append(sa, pair1626{
			score: scores[i],
			age:   ages[i],
		})
	}

	sort.Slice(sa, func(i, j int) bool {
		if sa[i].age == sa[j].age {
			return sa[i].score < sa[j].score
		}
		return sa[i].age < sa[j].age
	})

	dp := make([]int, len(scores))
	dp[0] = sa[0].score
	maxScore := dp[0]

	for i := 1; i < len(scores); i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			val := sa[i].score
			if sa[i].score >= sa[j].score {
				val += dp[j]
			}

			maxVal = Maximum(val, maxVal)
		}
		dp[i] = maxVal
		maxScore = Maximum(maxScore, dp[i])
	}

	return maxScore
}
