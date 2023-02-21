package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
			2,
		},
		{
			"test-2",
			args{[]int{3, 3, 7, 7, 10, 11, 11}},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNonDuplicate(tt.args.nums), "singleNonDuplicate(%v)", tt.args.nums)
		})
	}
}
