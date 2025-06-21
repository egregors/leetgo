package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				word: "aabcaba",
				k:    0,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				word: "dabdcbdcdcd",
				k:    2,
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				word: "aaabaaa",
				k:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumDeletions(tt.args.word, tt.args.k), "minimumDeletions(%v, %v)", tt.args.word, tt.args.k)
		})
	}
}
