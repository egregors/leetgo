package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "aab",
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				s: "aaabbbcc",
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				s: "ceabaacb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDeletions(tt.args.s), "minDeletions(%v)", tt.args.s)
		})
	}
}
