package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{word: "abbcccc"},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{word: "abcd"},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{word: "aaaa"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, possibleStringCount(tt.args.word), "possibleStringCount(%v)", tt.args.word)
		})
	}
}
