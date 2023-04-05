package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partitionString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{"abacaba"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, partitionString(tt.args.s), "partitionString(%v)", tt.args.s)
		})
	}
}
