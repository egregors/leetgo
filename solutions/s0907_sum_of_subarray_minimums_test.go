package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{3, 1, 2, 4},
			},
			17,
		},
		{
			"case 2",
			args{
				[]int{11, 81, 94, 43, 3},
			},
			444,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumSubarrayMins(tt.args.arr), "sumSubarrayMins(%v)", tt.args.arr)
		})
	}
}
