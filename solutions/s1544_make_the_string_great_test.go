package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "Example 2",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			name: "Example 3",
			args: args{
				s: "s",
			},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeGood(tt.args.s), "makeGood(%v)", tt.args.s)
		})
	}
}
