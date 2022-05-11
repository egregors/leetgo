package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countVowelStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{n: 1},
			5,
		},
		{
			"Test 1",
			args{n: 2},
			15,
		},
		{
			"Test 1",
			args{n: 33},
			66045,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countVowelStrings(tt.args.n), "countVowelStrings(%v)", tt.args.n)
		})
	}
}
