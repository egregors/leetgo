package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{
				"ABCABC",
				"ABC",
			},
			"ABC",
		},
		{
			"test-2",
			args{
				"ABABAB",
				"ABAB",
			},
			"AB",
		},
		{
			"test-3",
			args{
				"LEET",
				"CODE",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, gcdOfStrings(tt.args.str1, tt.args.str2), "gcdOfStrings(%v, %v)", tt.args.str1, tt.args.str2)
		})
	}
}
