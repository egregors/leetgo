package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_answerString(t *testing.T) {
	type args struct {
		word       string
		numFriends int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: word = "dbca", numFriends = 2
		//Output: "dbc"
		{
			name: "Test 1",
			args: args{
				word:       "dbca",
				numFriends: 2,
			},
			want: "dbc",
		},
		//Input: word = "gggg", numFriends = 4
		//Output: "g"
		{
			name: "Test 2",
			args: args{
				word:       "gggg",
				numFriends: 4,
			},
			want: "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, answerString(tt.args.word, tt.args.numFriends), "answerString(%v, %v)", tt.args.word, tt.args.numFriends)
		})
	}
}
