package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		//Input: words = ["mass","as","hero","superhero"]
		//Output: ["as","hero"]
		{
			name: "test1",
			args: args{
				words: []string{"mass", "as", "hero", "superhero"},
			},
			want: []string{"as", "hero"},
		},
		//Input: words = ["leetcode","et","code"]
		//Output: ["et","code"]
		{
			name: "test2",
			args: args{
				words: []string{"leetcode", "et", "code"},
			},
			want: []string{"et", "code"},
		},
		//Input: words = ["blue","green","bu"]
		//Output: []
		{
			name: "test3",
			args: args{
				words: []string{"blue", "green", "bu"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, stringMatching(tt.args.words), "stringMatching(%v)", tt.args.words)
		})
	}
}
