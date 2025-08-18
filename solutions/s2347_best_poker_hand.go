/*
	https://leetcode.com/problems/best-poker-hand

	You are given an integer array ranks and a character array suits. You have 5
		cards where the ith card has a rank of ranks[i] and a suit of suits[i].

	The following are the types of poker hands you can make from best to worst:

		"Flush": Five cards of the same suit.
		"Three of a Kind": Three cards of the same rank.
		"Pair": Two cards of the same rank.
		"High Card": Any single card.

	Return a string representing the best type of poker hand you can make with the
		given cards.

	Note that the return values are case-sensitive.
*/

package solutions

func bestHand(ranks []int, suits []byte) string {
	r := make(map[int]int)
	for _, rank := range ranks {
		r[rank]++
	}
	s := make(map[byte]bool)
	for _, suit := range suits {
		s[suit] = true
	}

	if len(s) == 1 {
		return "Flush"
	}

	for _, v := range r {
		if v >= 3 {
			return "Three of a Kind"
		}
	}

	for _, v := range r {
		if v == 2 {
			return "Pair"
		}
	}

	return "High Card"
}
