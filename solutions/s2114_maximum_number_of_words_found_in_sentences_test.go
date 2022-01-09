package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test #1",
			args: args{
				sentences: []string{"a b c", "b c", "c"},
			},
			want: 3,
		},
		{
			name: "test #1",
			args: args{
				sentences: []string{"a"},
			},
			want: 1,
		},
		{
			name: "test #1",
			args: args{
				sentences: []string{"a", "a b c", "a bb ccc d"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mostWordsFound(tt.args.sentences), "mostWordsFound(%v)", tt.args.sentences)
		})
	}
}
