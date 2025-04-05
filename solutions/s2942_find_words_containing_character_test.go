/*
	https://leetcode.com/problems/find-words-containing-character/description/

		You are given a 0-indexed array of strings words and a character x.

	Return an array of indices representing the words that contain the character x.

	Note that the returned array may be in any order.
*/

package solutions

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{words: []string{"leet", "code"}, x: byte('e')},
			[]int{0, 1},
		},
		{
			"test 2",
			args{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: byte('a')},
			[]int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
