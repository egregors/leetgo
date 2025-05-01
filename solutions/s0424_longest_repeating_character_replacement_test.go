package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, characterReplacement(tt.args.s, tt.args.k), "characterReplacement(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
