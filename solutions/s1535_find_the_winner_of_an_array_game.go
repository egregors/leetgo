/*
	https://leetcode.com/problems/find-the-winner-of-an-array-game/description/

	Given an integer array arr of distinct integers and an integer k.

	A game will be played between the first two elements of the array (i.e. arr[0]
		and arr[1]). In each round of the
	game, we compare arr[0] with arr[1], the larger integer wins and remains at
		position 0, and the smaller integer
	moves to the end of the array. The game ends when an integer wins k consecutive
		rounds.

	Return the integer which will win the game.

	It is guaranteed that there will be a winner of the game.
*/

//nolint:revive // it's ok
package solutions

type Buf1535 struct {
	Arr []int
	Cur int
}

func (b *Buf1535) inc() int     { return (b.Cur + 1) % len(b.Arr) }
func (b *Buf1535) Val() int     { return b.Arr[b.Cur] }
func (b *Buf1535) NextVal() int { return b.Arr[b.inc()] }
func (b *Buf1535) Inc()         { b.Cur = b.inc() }
func (b *Buf1535) Swap()        { b.Arr[b.Cur], b.Arr[b.inc()] = b.Arr[b.inc()], b.Arr[b.Cur] }

func getWinner(arr []int, k int) int {
	if k >= len(arr) {
		var maxEl int
		for _, x := range arr {
			if x > maxEl {
				maxEl = x
			}
		}

		return maxEl
	}

	var maxWins, winner, res int
	m := make(map[int]int)
	buf := Buf1535{Arr: arr}

	for maxWins < k {
		if buf.Val() > buf.NextVal() {
			winner = buf.Val()
			buf.Swap()
		} else {
			winner = buf.NextVal()
		}

		m[winner]++
		if m[winner] > maxWins {
			maxWins = m[winner]
			res = winner
		}

		buf.Inc()
	}

	return res
}
