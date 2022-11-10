package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates1047(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{s: "abbaca"},
			"ca",
		},
		{
			"Test 2",
			args{s: "azxxzy"},
			"ay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDuplicates1047(tt.args.s), "removeDuplicates1047(%v)", tt.args.s)
		})
	}
}
