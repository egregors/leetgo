package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent692(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			want: []string{"i", "love"},
		},
		{
			name: "test 2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			want: []string{"the", "is", "sunny", "day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topKFrequent692(tt.args.words, tt.args.k), "topKFrequent(%v, %v)", tt.args.words, tt.args.k)
		})
	}
}
