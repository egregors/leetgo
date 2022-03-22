package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				s: "aaabbb",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				s: "abab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkString(tt.args.s), "checkString(%v)", tt.args.s)
		})
	}
}
