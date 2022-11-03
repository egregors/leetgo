package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
		},
		{
			name: "Test 2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: [][]string(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(
				t,
				tt.want,
				findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList),
				"findLadders(%v, %v, %v)", tt.args.beginWord, tt.args.endWord, tt.args.wordList,
			)
		})
	}
}
