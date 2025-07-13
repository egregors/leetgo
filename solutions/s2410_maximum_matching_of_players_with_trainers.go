/*
	https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

	You are given a 0-indexed integer array players, where players[i] represents
		the ability of the ith player. You are also given a 0-indexed integer array
		trainers, where trainers[j] represents the training capacity of the jth
		trainer.

	The ith player can match with the jth trainer if the player's ability is less
		than or equal to the trainer's training capacity. Additionally, the ith player
		can be matched with at most one trainer, and the jth trainer can be matched
		with at most one player.

	Return the maximum number of matchings between players and trainers that
		satisfy these conditions.
*/

package solutions

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) (ans int) {
	sort.Ints(players)
	sort.Ints(trainers)
	m, n := len(players), len(trainers)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && players[i] > trainers[j] {
			j++
		}
		if j < n {
			ans++
			j++
		}
	}
	return
}
