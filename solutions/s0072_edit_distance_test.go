package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDistance72(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{"horse", "ros"},
			3,
		},
		{
			"test 2",
			args{"intention", "execution"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDistance72(tt.args.word1, tt.args.word2), "minDistance72(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
