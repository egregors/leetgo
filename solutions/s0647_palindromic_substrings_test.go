package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test 1",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubstrings(tt.args.s), "countSubstrings(%v)", tt.args.s)
		})
	}
}
