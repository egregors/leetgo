package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case 1",
			args{[]string{"01", "10"}},
			"11",
		},
		{
			"case 2",
			args{[]string{"00", "01"}},
			"10",
		},
		{
			"case 3",
			args{[]string{"111", "011", "001"}},
			"000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findDifferentBinaryString(tt.args.nums), "findDifferentBinaryString(%v)", tt.args.nums)
		})
	}
}
