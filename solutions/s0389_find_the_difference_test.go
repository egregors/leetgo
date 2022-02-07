package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example 1",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e',
		},
		{
			name: "Example 2",
			args: args{
				s: "",
				t: "a",
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTheDifference(tt.args.s, tt.args.t), "findTheDifference(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
