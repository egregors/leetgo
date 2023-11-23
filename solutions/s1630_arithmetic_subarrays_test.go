package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	type args struct {
		nums []int
		l    []int
		r    []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case 0",
			args: args{
				nums: []int{4, 6, 5, 9, 3, 7},
				l:    []int{0, 0, 2},
				r:    []int{2, 3, 5},
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkArithmeticSubarrays(tt.args.nums, tt.args.l, tt.args.r), "checkArithmeticSubarrays(%v, %v, %v)", tt.args.nums, tt.args.l, tt.args.r)
		})
	}
}
