package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[]int{1, 2, 1}},
			3,
		},
		{
			"test 2",
			args{[]int{0, 1, 2, 2}},
			3,
		},
		{
			"test 3",
			args{[]int{1, 2, 3, 2, 2}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, totalFruit(tt.args.fruits), "totalFruit(%v)", tt.args.fruits)
		})
	}
}
