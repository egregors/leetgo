/*
	https://leetcode.com/problems/make-the-string-great/

		Given a string s of lower and upper case English letters.

	A good string is a string which doesn't have two adjacent characters s[i] and
		s[i + 1]
	where:

		0 <= i <= s.length - 2
		s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case
			or
	vice-versa.

	To make the string good, you can choose two adjacent characters that make the
		string
	bad and remove them. You can keep doing this until the string becomes good.

	Return the string after making it good. The answer is guaranteed to be unique
		under
	the given constraints.

	Notice that an empty string is also good.
*/

//nolint:revive // it's ok
package solutions

type Stack1544 []rune

func (s *Stack1544) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack1544) Push(r rune)   { *s = append(*s, r) }
func (s *Stack1544) Peek() rune    { return (*s)[len(*s)-1] }

func (s *Stack1544) Pop() rune {
	r := s.Peek()
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack1544) ToSlice() []rune {
	return *s
}

func makeGood(s string) string {
	st := Stack1544{}
	for _, r := range s {
		if !st.IsEmpty() && Abs(int(st.Peek()-r)) == 32 {
			st.Pop()
		} else {
			st.Push(r)
		}
	}

	return string(st.ToSlice())
}
