package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestNumber(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{"IIIDIDDD"},
			"123549876",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, smallestNumber(tt.args.pattern), "smallestNumber(%v)", tt.args.pattern)
		})
	}
}
