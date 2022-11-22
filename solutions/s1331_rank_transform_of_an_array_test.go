package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[]int{40, 10, 20, 30}},
			[]int{4, 1, 2, 3},
		},
		{
			"case 2",
			args{[]int{100, 100, 100}},
			[]int{1, 1, 1},
		},
		{
			"case 3",
			args{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}},
			[]int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, arrayRankTransform(tt.args.arr), "arrayRankTransform(%v)", tt.args.arr)
		})
	}
}
