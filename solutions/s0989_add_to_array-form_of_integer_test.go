package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				[]int{1, 2, 0, 0},
				34,
			},
			[]int{1, 2, 3, 4},
		},
		{
			"test 2",
			args{
				[]int{2, 7, 4},
				181,
			},
			[]int{4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addToArrayForm(tt.args.num, tt.args.k), "addToArrayForm(%v, %v)", tt.args.num, tt.args.k)
		})
	}
}
