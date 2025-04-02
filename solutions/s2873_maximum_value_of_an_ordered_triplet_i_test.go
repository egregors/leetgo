package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"case 1",
			args{[]int{12, 6, 1, 2, 7}},
			77,
		},
		{
			"case 2",
			args{[]int{1, 10, 3, 4, 19}},
			133,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumTripletValue(tt.args.nums), "maximumTripletValue(%v)", tt.args.nums)
		})
	}
}
