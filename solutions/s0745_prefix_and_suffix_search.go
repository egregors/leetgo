/*
	https://leetcode.com/problems/prefix-and-suffix-search/

	Design a special dictionary with some words that searches the words in it by a
		prefix
	and a suffix.

	Implement the WordFilter class:

		WordFilter(string[] words) Initializes the object with the words in the
			dictionary.
		f(string prefix, string suffix) Returns the index of the word in the
			dictionary, which
			has the prefix prefix and the suffix suffix. If there is more than one valid
				index,
			return the largest of them. If there is no such word in the dictionary,
				return -1.
*/

//nolint:revive // it's ok
package solutions

type TrieNode745 struct {
	children []*TrieNode745
	weight   int
}

func NewTrieNode() *TrieNode745 {
	return &TrieNode745{
		children: make([]*TrieNode745, 27),
		weight:   0,
	}
}

type WordFilter struct {
	trie *TrieNode745
}

// NewWordFilter should call Constructor to pass LeetCode tests
func NewWordFilter(words []string) WordFilter {
	wf := WordFilter{trie: NewTrieNode()}
	for weight := 0; weight < len(words); weight++ {
		word := words[weight] + "{"
		for i := 0; i < len(word); i++ {
			cur := wf.trie
			cur.weight = weight
			for j := i; j < 2*len(word)-1; j++ {
				k := word[j%len(word)] - 'a'
				if cur.children[k] == nil {
					cur.children[k] = NewTrieNode()
				}
				cur = cur.children[k]
				cur.weight = weight
			}
		}
	}
	return wf
}

func (wf *WordFilter) F(prefix, suffix string) int {
	cur := wf.trie
	for _, letter := range suffix + "{" + prefix {
		if cur.children[letter-'a'] == nil {
			return -1
		}
		cur = cur.children[letter-'a']
	}
	return cur.weight
}
