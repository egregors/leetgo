package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates1209(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 0",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "abcd",
		},
		{
			name: "Test 1",
			args: args{
				s: "deeedbbcccbdaa",
				k: 3,
			},
			want: "aa",
		},
		{
			name: "Test 2",
			args: args{
				s: "pbbcggttciiippooaais",
				k: 2,
			},
			want: "ps",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDuplicates1209(tt.args.s, tt.args.k), "removeDuplicates1209(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
