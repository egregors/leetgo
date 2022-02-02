package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
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
				nums: []int{1, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1},
			},
			want: 5,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
			},
			want: 8,
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1},
			},
			want: 8,
		},
		{
			name: "Test 7",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findPeakElement(tt.args.nums), "findPeakElement(%v)", tt.args.nums)
		})
	}
}
