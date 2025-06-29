package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validWordAbbreviation(t *testing.T) {
	type args struct {
		word string
		abbr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//Example 1:
		//
		//Input: word = "internationalization", abbr = "i12iz4n"
		//Output: true
		//Explanation: The word "internationalization" can be abbreviated as "i12iz4n" ("i nternational iz atio n").
		//
		//Example 2:
		//
		//Input: word = "apple", abbr = "a2e"
		//Output: false
		//Explanation: The word "apple" cannot be abbreviated as "a2e".
		{
			name: "Example 1",
			args: args{
				word: "internationalization",
				abbr: "i12iz4n",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				word: "apple",
				abbr: "a2e",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validWordAbbreviation(tt.args.word, tt.args.abbr), "validWordAbbreviation(%v, %v)", tt.args.word, tt.args.abbr)
		})
	}
}
