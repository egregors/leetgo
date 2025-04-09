package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				nums: []int{4, 5, 6, 4, 4},
			},
			want: 2,
		},
		{
			name: "Test3",
			args: args{
				nums: []int{6, 7, 8, 9},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumOperations(tt.args.nums), "minimumOperations(%v)", tt.args.nums)
		})
	}
}
