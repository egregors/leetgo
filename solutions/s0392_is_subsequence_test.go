package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isSubsequence(tt.args.s, tt.args.t), "isSubsequence(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
