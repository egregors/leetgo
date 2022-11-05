package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"eat", "oath"},
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"abcb"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, findWords(tt.args.board, tt.args.words), "findWords(%v, %v)", tt.args.board, tt.args.words)
		})
	}
}
