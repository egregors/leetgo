package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteAndEarn(t *testing.T) {
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
			args: args{[]int{3, 4, 2}},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{[]int{2, 2, 3, 3, 3, 4}},
			want: 9,
		},
		{
			name: "Test 3",
			args: args{[]int{2, 2, 3, 3, 3, 4}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deleteAndEarn(tt.args.nums), "deleteAndEarn(%v)", tt.args.nums)
		})
	}
}
