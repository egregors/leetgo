package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mincostTickets(tt.args.days, tt.args.costs), "mincostTickets(%v, %v)", tt.args.days, tt.args.costs)
		})
	}
}
