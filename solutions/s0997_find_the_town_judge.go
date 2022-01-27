/*
	https://leetcode.com/problems/find-the-town-judge/

	In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

	If the town judge exists, then:

		The town judge trusts nobody.
		Everybody (except for the town judge) trusts the town judge.
		There is exactly one person that satisfies properties 1 and 2.

	You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai
	trusts the person labeled bi.

	Return the label of the town judge if the town judge exists and can be identified, or return -1
	otherwise.
*/

package solutions

func contains(xs []int, k int) bool {
	for _, x := range xs {
		if x == k {
			return true
		}
	}
	return false
}

func findJudge(n int, trust [][]int) int {
	m := make(map[int][]int, n)
	for i := 1; i <= n; i++ {
		m[i] = make([]int, 0)
	}
	for _, t := range trust {
		a, b := t[0], t[1]
		m[a] = append(m[a], b)
	}
	for k, v := range m {
		if len(v) == 0 {

			i := 0
			for _, v := range m {
				if contains(v, k) {
					i++
				}
			}

			if i == len(m)-1 {
				return k
			}
		}
	}

	return -1
}
