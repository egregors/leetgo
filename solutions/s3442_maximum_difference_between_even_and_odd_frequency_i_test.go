package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDifference(t *testing.T) {
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
			args: args{s: "aaaaabbc"},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{s: "abcabcab"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxDifference(tt.args.s), "maxDifference(%v)", tt.args.s)
		})
	}
}
