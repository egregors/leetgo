package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				s: "a1b2",
			},
			want: []string{
				"A1B2", "a1B2", "A1b2", "a1b2",
			},
		},
		{
			name: "Example 2",
			args: args{
				s: "8gdG",
			},
			want: []string{
				"8GDG", "8gDG", "8GdG", "8gdG", "8GDg", "8gDg", "8Gdg", "8gdg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, letterCasePermutation(tt.args.s), "letterCasePermutation(%v)", tt.args.s)
		})
	}
}
