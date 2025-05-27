/*
	We are playing the Guess Game. The game is as follows:

	I pick a number from 1 to n. You have to guess which number I picked.

	Every time you guess wrong, I will tell you whether the number I picked is
		higher or lower than your guess.

	You call a pre-defined API int guess(int num), which returns three possible
		results:

		-1: Your guess is higher than the number I picked (i.e. num > pick).
		1: Your guess is lower than the number I picked (i.e. num < pick).
		0: your guess is equal to the number I picked (i.e. num == pick).

	Return the number that I picked.
*/

package solutions

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func makeGuess(pick int) func(num int) int {
	return func(num int) int {
		if num < pick {
			return 1
		}
		if num > pick {
			return -1
		}
		return 0
	}
}

// guessNumber shouldn't have `guess` argument to pass LeetCode test. need it for tests
func guessNumber(n int, guess func(int) int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		}
	}
	return -1
}
