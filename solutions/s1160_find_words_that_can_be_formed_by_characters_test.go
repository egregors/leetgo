/*
	https://leetcode.com/problems/find-words-that-can-be-formed-by-characters

	You are given an array of strings words and a string chars.

	A string is good if it can be formed by characters from chars (each character can only
	be used once for each word in words).

	Return the sum of lengths of all good strings in words.
*/

package solutions

import "testing"

func Test_countCharacters(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				words: []string{"cat", "bt", "hat", "tree"},
				chars: "atach",
			},
			6,
		},
		{
			"test2",
			args{
				words: []string{"hello", "world", "leetcode"},
				chars: "welldonehoneyr",
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharacters(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("countCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
