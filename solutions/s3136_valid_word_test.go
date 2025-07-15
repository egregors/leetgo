package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid3136(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"234Adas"},
			true,
		},
		{
			"test2",
			args{"b3"},
			false,
		},
		{
			"test3",
			args{"a3$e"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValid3136(tt.args.word), "isValid3136(%v)", tt.args.word)
		})
	}
}
