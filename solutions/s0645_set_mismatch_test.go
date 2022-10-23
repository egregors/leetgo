package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-1",
			args: args{
				nums: []int{1, 2, 2, 4},
			},
			want: []int{2, 3},
		},
		{
			name: "test-2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findErrorNums(tt.args.nums), "findErrorNums(%v)", tt.args.nums)
		})
	}
}
