package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_closeStrings(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{
				"abc",
				"bca",
			},
			true,
		},
		{
			"test-2",
			args{
				"a",
				"aa",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, closeStrings(tt.args.word1, tt.args.word2), "closeStrings(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
