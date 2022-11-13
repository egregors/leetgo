package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords151(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "2",
			args: args{
				s: "  hello world!  ",
			},
			want: "world! hello",
		},
		{
			name: "3",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseWords151(tt.args.s), "reverseWords151(%v)", tt.args.s)
		})
	}
}
