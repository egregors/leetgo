package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{
					"gin", "zen", "gig", "msg",
				},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				words: []string{
					"a",
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, uniqueMorseRepresentations(tt.args.words), "uniqueMorseRepresentations(%v)", tt.args.words)
		})
	}
}
