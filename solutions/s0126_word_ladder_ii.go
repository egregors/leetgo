/*
	https://leetcode.com/problems/word-ladder-ii/

	A transformation sequence from word beginWord to word endWord using a
	dictionary wordList is a sequence of words
	beginWord -> s1 -> s2 -> ... -> sk such that:

		Every adjacent pair of words differs by a single letter.
		Every si for 1 <= i <= k is in wordList. Note that beginWord does not
	need to be in wordList.
		sk == endWord

	Given two words, beginWord and endWord, and a dictionary wordList, return
	all the shortest transformation sequences from beginWord to endWord, or an
	empty list if no such sequence exists. Each sequence should be returned as
	a list of the words [beginWord, s1, s2, ..., sk].
*/

package solutions

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]struct{})
	current, trace := make(map[string]struct{}), make(map[string][]string)
	for _, s := range wordList {
		dict[s] = struct{}{}
		trace[s] = make([]string, 0)
	}
	dict[beginWord] = struct{}{}
	trace[beginWord] = make([]string, 0)
	current[beginWord] = struct{}{}
	_, ok := current[endWord]
	for len(current) != 0 && !ok {
		for word := range current {
			delete(dict, word)
		}
		next := make(map[string]struct{})
		for word := range current {
			for i := range word {
				for _, c := range "abcdefghijklmnopqrstuvwxyz" {
					candidate := word[:i] + string(c) + word[i+1:]
					if _, getOk := dict[candidate]; getOk {
						trace[candidate] = append(trace[candidate], word)
						next[candidate] = struct{}{}
					}
				}
			}
		}
		current = next
		_, ok = current[endWord]
	}
	var results [][]string
	if len(current) != 0 {
		backtrace(&results, trace, []string{}, endWord)
	}
	return results
}

func backtrace(results *[][]string, trace map[string][]string, path []string, word string) {
	if len(trace[word]) == 0 {
		*results = append(*results, append([]string{word}, path...))
	} else {
		for _, prev := range trace[word] {
			backtrace(results, trace, append([]string{word}, path...), prev)
		}
	}
}
