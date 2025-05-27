/*
	You are given a string s and an integer k, a k duplicate removal consists of
	choosing k adjacent and equal letters from s and removing them, causing the
		left
	and the right side of the deleted substring to concatenate together.

	We repeatedly make k duplicate removals on s until we no longer can.

	Return the final string after all such duplicate removals have been made. It is
	guaranteed that the answer is unique.
*/

package solutions

type stack1209 []int

func (s *stack1209) Push(x int) { *s = append(*s, x) }
func (s *stack1209) Peek() int  { return (*s)[len(*s)-1] }
func (s *stack1209) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// removeDuplicates1209 should call removeDuplicates to pass LeetCode tests
func removeDuplicates1209(s string, k int) string {
	counts := stack1209{}
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(counts) == 0 || rs[i-1] != rs[i] {
			counts.Push(1)
		} else {
			inc := counts.Pop() + 1
			if inc == k {
				rs = append(rs[:i-k+1], rs[i+1:]...)
				i -= k
			} else {
				counts.Push(inc)
			}
		}
	}
	return string(rs)
}
