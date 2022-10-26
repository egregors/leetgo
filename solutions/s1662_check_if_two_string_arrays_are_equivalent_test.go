package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayStringsAreEqual(t *testing.T) {
	type args struct {
		word1 []string
		word2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				word1: []string{"ab", "c"},
				word2: []string{"a", "bc"},
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				word1: []string{"a", "cb"},
				word2: []string{"ab", "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, arrayStringsAreEqual(tt.args.word1, tt.args.word2), "arrayStringsAreEqual(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
