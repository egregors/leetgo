package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{[]int{1, 2, 3, 5, 7}},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfArithmeticSlices(tt.args.nums), "numberOfArithmeticSlices(%v)", tt.args.nums)
		})
	}
}
