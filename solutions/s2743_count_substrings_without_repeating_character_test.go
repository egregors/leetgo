package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfSpecialSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "abcd",
			},
			want: 10,
		},
		{
			name: "Test 2",
			args: args{
				s: "aaa",
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				s: "abab",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfSpecialSubstrings(tt.args.s), "numberOfSpecialSubstrings(%v)", tt.args.s)
		})
	}
}
