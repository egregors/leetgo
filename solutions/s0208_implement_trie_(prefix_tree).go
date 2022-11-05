/*
	https://leetcode.com/problems/implement-trie-prefix-tree/

	A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently
	store and retrieve keys in a dataset of strings. There are various applications of this
	data structure, such as autocomplete and spellchecker.

	Implement the Trie class:

		Trie() Initializes the trie object.
		void insert(String word) Inserts the string word into the trie.
		boolean search(String word) Returns true if the string word is in the trie (i.e., was
	inserted before), and false otherwise.
		boolean startsWith(String prefix) Returns true if there is a previously inserted string
	word that has the prefix prefix, and false otherwise.
*/

//nolint:revive // it's ok
package solutions

type Trie struct {
	ch       int32
	key      bool
	children [26]*Trie
}

func Constructor() Trie {
	return *NewTrie(0)
}

func NewTrie(ch int32) *Trie {
	return &Trie{ch: ch, children: [26]*Trie{}}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, r := range word {
		ch := r - 'a'
		if node.children[ch] == nil {
			node.children[ch] = NewTrie(ch)
		}
		node = node.children[ch]
	}
	node.key = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, r := range word {
		ch := r - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.key
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, r := range prefix {
		ch := r - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node != nil
}
