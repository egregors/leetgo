package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canBeTypedWords(t *testing.T) {
	type args struct {
		text          string
		brokenLetters string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				text:          "hello world",
				brokenLetters: "ad",
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				text:          "leet code",
				brokenLetters: "lt",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				text:          "leet code",
				brokenLetters: "e",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canBeTypedWords(tt.args.text, tt.args.brokenLetters), "canBeTypedWords(%v, %v)", tt.args.text, tt.args.brokenLetters)
		})
	}
}
