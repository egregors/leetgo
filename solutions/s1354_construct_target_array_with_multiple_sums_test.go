package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPossible(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 0",
			args{target: []int{9, 3, 5}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPossible(tt.args.target), "isPossible(%v)", tt.args.target)
		})
	}
}
