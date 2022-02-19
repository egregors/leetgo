package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantCounter int
	}{
		{
			"Test 1",
			args{s: "xyzzaz"},
			1,
		},
		{
			"Test 2",
			args{s: "aababcabc"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantCounter, countGoodSubstrings(tt.args.s), "countGoodSubstrings(%v)", tt.args.s)
		})
	}
}
