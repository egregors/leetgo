package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, summaryRanges(tt.args.nums), "summaryRanges(%v)", tt.args.nums)
		})
	}
}
