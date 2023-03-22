/*
	https://leetcode.com/problems/minimum-genetic-mutation/

	A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.

	Suppose we need to investigate a mutation from a gene string start to a gene string end where one mutation is
	defined as one single character changed in the gene string.

		For example, "AACCGGTT" --> "AACCGGTA" is one mutation.

	There is also a gene bank bank that records all the valid gene mutations. A gene must be in bank to make it a
	valid gene string.

	Given the two gene strings start and end and the gene bank bank, return the minimum number of mutations needed
	to mutate from start to end. If there is no such a mutation, return -1.

	Note that the starting point is assumed to be valid, so it might not be included in the bank.
*/

//nolint:revive // it's ok
package solutions

type Queue433 []string

func (q *Queue433) Push(s string) { *q = append(*q, s) }
func (q *Queue433) IsEmpty() bool { return len(*q) == 0 }
func (q *Queue433) Len() int      { return len(*q) }
func (q *Queue433) Pop() string {
	if !q.IsEmpty() {
		el := (*q)[0]
		*q = (*q)[1:]
		return el
	}
	return ""
}

func minMutation(start, end string, bank []string) int {
	q := Queue433{}
	seen := make(map[string]struct{})

	q.Push(start)
	seen[start] = struct{}{}

	steps := 0

	for !q.IsEmpty() {
		nsInQ := q.Len()
		for i := 0; i < nsInQ; i++ {
			n := q.Pop()
			if n == end {
				return steps
			}
			for _, ch := range []rune{'A', 'C', 'G', 'T'} {
				for i := 0; i < len(n); i++ {
					neighbor := n[:i] + string(ch) + n[i+1:]
					if _, ok := seen[neighbor]; !ok && Contains(bank, neighbor) {
						q.Push(neighbor)
						seen[neighbor] = struct{}{}
					}
				}
			}
		}
		steps++
	}

	return -1
}
