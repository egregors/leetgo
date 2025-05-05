package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canPermutePalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "code",
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				s: "aab",
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				s: "carerac",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPermutePalindrome(tt.args.s), "canPermutePalindrome(%v)", tt.args.s)
		})
	}
}
