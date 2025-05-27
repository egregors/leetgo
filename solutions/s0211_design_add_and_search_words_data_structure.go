/*
	https://leetcode.com/problems/design-add-and-search-words-data-structure/

	Design a data structure that supports adding new words and finding if a string
		matches any previously added
	string.

	Implement the WordDictionary class:

		WordDictionary() Initializes the object.
		void addWord(word) Adds word to the data structure, it can be matched later.
		bool search(word) Returns true if there is any string in the data structure
			that matches word or false
	otherwise. word may contain dots '.' where dots can be matched with any letter.

*/

//nolint:revive //it's ok
package solutions

type TrieNode211 struct {
	IsWord bool
	Next   map[rune]*TrieNode211
}

func NewTrieNode211() *TrieNode211 {
	return &TrieNode211{
		IsWord: false,
		Next:   make(map[rune]*TrieNode211),
	}
}

func (tn *TrieNode211) AddWord(word string) {
	if word == "" {
		tn.IsWord = true
		return
	}
	if _, ok := tn.Next[rune(word[0])]; !ok {
		tn.Next[rune(word[0])] = NewTrieNode211()
	}
	tn.Next[rune(word[0])].AddWord(word[1:])
}

func (tn *TrieNode211) Search(word string) bool {
	if word == "" {
		return tn.IsWord
	}
	if word[0] == '.' {
		for _, v := range tn.Next {
			if v.Search(word[1:]) {
				return true
			}
		}
		return false
	}
	if _, ok := tn.Next[rune(word[0])]; !ok {
		return false
	}
	return tn.Next[rune(word[0])].Search(word[1:])
}

type WordDictionary struct {
	trie *TrieNode211
}

// NewWordDictionary should call Constructor to pass LeeCode testa
func NewWordDictionary() WordDictionary {
	return WordDictionary{
		trie: &TrieNode211{Next: make(map[rune]*TrieNode211)},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	wd.trie.AddWord(word)
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.trie.Search(word)
}
