/*
	Given an array of n integers nums, a 132 pattern is a subsequence of three
		integers nums[i], nums[j] and nums[k]
	such that i < j < k and nums[i] < nums[k] < nums[j].

	Return true if there is a 132 pattern in nums, otherwise, return false.
*/

package solutions

type stack456 []int

func (s *stack456) IsEmpty() bool { return len(*s) == 0 }
func (s *stack456) Peek() int     { return (*s)[len(*s)-1] }
func (s *stack456) Push(x int)    { *s = append(*s, x) }
func (s *stack456) Pop() int {
	x := s.Peek()
	*s = (*s)[:len(*s)-1]
	return x
}

func find132pattern(nums []int) bool {
	stack := stack456{}
	min := make([]int, len(nums))
	min[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		min[i] = Minimum(min[i-1], nums[i])
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] > min[j] {
			for !stack.IsEmpty() && stack.Peek() <= min[j] {
				stack.Pop()
			}
			if !stack.IsEmpty() && stack.Peek() < nums[j] {
				return true
			}
			stack.Push(nums[j])
		}
	}
	return false
}
