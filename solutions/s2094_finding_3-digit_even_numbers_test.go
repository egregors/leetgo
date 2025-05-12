package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findEvenNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{digits: []int{2, 1, 3, 0}},
			[]int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{
			"Test 2",
			args{digits: []int{2, 2, 8, 8, 2}},
			[]int{222, 228, 282, 288, 822, 828, 882},
		},
		{
			"Test 3",
			args{digits: []int{3, 7, 5}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findEvenNumbers(tt.args.digits), "findEvenNumbers(%v)", tt.args.digits)
		})
	}
}
