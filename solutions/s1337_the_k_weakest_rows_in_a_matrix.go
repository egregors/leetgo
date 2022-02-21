/*
	https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

	You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's
	(representing civilians). The soldiers are positioned in front of the civilians.
	That is, all the 1's will appear to the left of all the 0's in each row.

	A row i is weaker than a row j if one of the following is true:

		The number of soldiers in row i is less than the number of soldiers in row j.
		Both rows have the same number of soldiers and i < j.

	Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.
*/

package solutions

func kWeakestRows(mat [][]int, k int) []int { //nolint:revive // this is LeenCode name
	var res []int
	seen := make(map[int]bool)

	for c := 0; c < len(mat[0]); c++ {
		for r := 0; r < len(mat); r++ {
			if !seen[r] {
				if mat[r][c] == 0 {
					seen[r] = true
					res = append(res, r)
					if len(res) == k {
						return res
					}
				}
			}
		}
	}

	for i := 0; i < len(mat) && len(res) < k; i++ {
		if mat[i][len(mat[0])-1] == 1 {
			res = append(res, i)
		}
	}

	return res
}
