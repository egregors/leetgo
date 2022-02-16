package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_thirdMax(t *testing.T) {
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
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, thirdMax(tt.args.nums), "thirdMax(%v)", tt.args.nums)
		})
	}
}
