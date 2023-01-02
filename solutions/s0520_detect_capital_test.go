package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{"USA"},
			true,
		},
		{
			"case 2",
			args{"FlaG"},
			false,
		},
		{
			"case 3",
			args{"leetcode"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, detectCapitalUse(tt.args.word), "detectCapitalUse(%v)", tt.args.word)
		})
	}
}
