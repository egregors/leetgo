package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{3, 7},
			3,
		},
		{
			"Example 2",
			args{8, 10},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countOdds(tt.args.low, tt.args.high), "countOdds(%v, %v)", tt.args.low, tt.args.high)
		})
	}
}
