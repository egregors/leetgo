package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				"00110",
			},
			1,
		},
		{
			"case 2",
			args{
				"010110",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minFlipsMonoIncr(tt.args.s), "minFlipsMonoIncr(%v)", tt.args.s)
		})
	}
}
