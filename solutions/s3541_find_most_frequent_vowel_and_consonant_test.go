package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFreqSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: "successes"},
			want: 6,
		},
		{
			name: "test2",
			args: args{s: "aeiaeia"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxFreqSum(tt.args.s), "maxFreqSum(%v)", tt.args.s)
		})
	}
}
