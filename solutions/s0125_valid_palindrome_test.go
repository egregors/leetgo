package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome125(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "Test2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "Test3",
			args: args{
				s: " ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPalindrome125(tt.args.s), "isPalindrome125(%v)", tt.args.s)
		})
	}
}
