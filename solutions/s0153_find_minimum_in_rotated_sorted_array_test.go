package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {
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
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 1,
		},
		{
			name: "Test 7",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMin(tt.args.nums), "findMin(%v)", tt.args.nums)
		})
	}
}
