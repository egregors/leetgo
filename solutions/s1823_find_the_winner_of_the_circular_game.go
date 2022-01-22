/*
	https://leetcode.com/problems/find-the-winner-of-the-circular-game/

	There are n friends that are playing a game. The friends are sitting in a circle and are numbered
	from 1 to n in clockwise order. More formally, moving clockwise from the ith friend brings you to
	the (i+1)th friend for 1 <= i < n, and moving clockwise from the nth friend brings you to the 1st friend.

	The rules of the game are as follows:

		Start at the 1st friend.
		Count the next k friends in the clockwise direction including the friend you started at. The counting
		wraps around the circle and may count some friends more than once.
		The last friend you counted leaves the circle and loses the game.
		If there is still more than one friend in the circle, go back to step 2 starting from the friend
		immediately clockwise of the friend who just lost and repeat.
		Else, the last friend in the circle wins the game.

	Given the number of friends, n, and an integer k, return the winner of the game.
*/

package solutions

func roll(ps []int, from, k int) []int {
	if len(ps) == 1 {
		return ps
	}
	toDel := (from + k - 1) % len(ps)
	RemoveFromIntSlice(&ps, toDel)
	return roll(ps, toDel%len(ps), k)
}

func findTheWinner(n, k int) int {
	var ps []int
	for i := 1; i <= n; i++ {
		ps = append(ps, i)
	}
	return roll(ps, 0, k)[0]
}
