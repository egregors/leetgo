/*
	https://leetcode.com/problems/delete-columns-to-make-sorted/description/

	You are given an array of n strings strs, all of the same length.

	The strings can be arranged such that there is one on each line, making a grid. For example,
	strs = ["abc", "bce", "cae"] can be arranged as:

	abc
	bce
	cae

	You want to delete the columns that are not sorted lexicographically. In the above example
	(0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted while column 1
	('b', 'c', 'a') is not, so you would delete column 1.

	Return the number of columns that you will delete.
*/

package solutions

func minDeletionSize(strs []string) int {
	deleted := 0
	for c := 0; c < len(strs[0]); c++ {
		for r := 0; r < len(strs)-1; r++ {
			if strs[r][c] > strs[r+1][c] {
				deleted++
				break
			}
		}
	}
	return deleted
}
