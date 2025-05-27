/*
	https://leetcode.com/problems/search-suggestions-system/

	You are given an array of strings products and a string searchWord.

	Design a system that suggests at most three product names from products after
		each character
	of searchWord is typed. Suggested products should have common prefix with
		searchWord.
	If there are more than three products with a common prefix return the three
		lexicographically
	minimums products.

	Return a list of lists of the suggested products after each character of
		searchWord is typed.
*/

//nolint:revive // it's ok
package solutions

import (
	"fmt"
	"sort"
)

type TrieNode1268 struct {
	children [26]*TrieNode1268
	isWord   bool
}

func (n *TrieNode1268) String() string {
	var w string
	for i, c := range n.children {
		if c != nil {
			w += string(rune(i + 'a'))
		}
	}
	return fmt.Sprintf("ch: [%s]", w)
}

func (n *TrieNode1268) Insert(s string) {
	curr := n
	for _, c := range s {
		if curr.children[c-'a'] == nil {
			curr.children[c-'a'] = &TrieNode1268{}
		}
		curr = curr.children[c-'a']
	}
	curr.isWord = true
}

func (n *TrieNode1268) GetWordsStartingWith(prefix string) []string {
	curr := n
	var resBuf []string
	for _, c := range prefix {
		if curr.children[c-'a'] == nil {
			return resBuf
		}
		curr = curr.children[c-'a']
	}
	curr.DFSWithPrefix(prefix, &resBuf)
	return resBuf
}

func (n *TrieNode1268) DFSWithPrefix(prefix string, acc *[]string) {
	if len(*acc) == 3 {
		return
	}
	if n.isWord {
		*acc = append(*acc, prefix)
	}

	for c := 'a'; c <= 'z'; c++ {
		if n.children[c-'a'] != nil {
			n.children[c-'a'].DFSWithPrefix(prefix+string(c), acc)
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	root := &TrieNode1268{}
	for _, p := range products {
		root.Insert(p)
	}

	var prefix string
	var res [][]string //nolint:prealloc // size is unknown

	for _, c := range searchWord {
		prefix += string(c)
		res = append(res, root.GetWordsStartingWith(prefix))
	}

	return res
}
