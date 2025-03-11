package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 0",
			args{"abcabc"},
			10,
		},
		{
			"case 1",
			args{"aaacb"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfSubstrings(tt.args.s), "numberOfSubstrings(%v)", tt.args.s)
		})
	}
}
