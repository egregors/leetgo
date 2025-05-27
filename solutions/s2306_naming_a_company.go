/*
	https://leetcode.com/problems/naming-a-company/

	You are given an array of strings ideas that represents a list of names to be
		used in
	the process of naming a company. The process of naming a company is as follows:

    Choose 2 distinct names from ideas, call them ideaA and ideaB.
    Swap the first letters of ideaA and ideaB with each other.
    If both of the new names are not found in the original ideas, then the name
    	ideaA ideaB
	(the concatenation of ideaA and ideaB, separated by a space) is a valid company
		name.
    Otherwise, it is not a valid name.

	Return the number of distinct valid names for the company.
*/

package solutions

type set2306 map[string]struct{}

func (s set2306) add(str string) {
	s[str] = struct{}{}
}

func (s set2306) contains(str string) bool {
	_, ok := s[str]
	return ok
}

func distinctNames(ideas []string) int64 {
	initialSet := make([]set2306, 26)
	for i := 0; i < 26; i++ {
		initialSet[i] = make(set2306)
	}

	for _, idea := range ideas {
		initialSet[idea[0]-'a'].add(idea[1:])
	}

	ans := 0
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			numOfMutual := 0
			for ideaA := range initialSet[i] {
				if initialSet[j].contains(ideaA) {
					numOfMutual++
				}
			}
			ans += 2 * (len(initialSet[i]) - numOfMutual) * (len(initialSet[j]) - numOfMutual)
		}
	}
	return int64(ans)
}
