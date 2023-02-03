package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case 1",
			args{
				"PAYPALISHIRING",
				3,
			},
			"PAHNAPLSIIGYIR",
		},
		{
			"case 2",
			args{
				"PAYPALISHIRING",
				4,
			},
			"PINALSIGYAHRPI",
		},
		{
			"case 3",
			args{
				"A",
				1,
			},
			"A",
		},
		{
			"case 4",
			args{
				"AB",
				1,
			},
			"AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convert(tt.args.s, tt.args.numRows), "convert(%v, %v)", tt.args.s, tt.args.numRows)
		})
	}
}
