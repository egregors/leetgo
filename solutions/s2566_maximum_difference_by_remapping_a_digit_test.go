package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMaxDifference(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{num: 11891},
			want: 99009,
		},
		{
			name: "Example 2",
			args: args{num: 90},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minMaxDifference(tt.args.num), "minMaxDifference(%v)", tt.args.num)
		})
	}
}
