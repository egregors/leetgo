package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sum2235(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				num1: 2,
				num2: 3,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sum2235(tt.args.num1, tt.args.num2), "sum2235(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
