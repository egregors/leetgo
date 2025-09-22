package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFrequencyElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 2, 2, 3, 1, 4}},
			want: 4,
		},
		{
			name: "test2",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxFrequencyElements(tt.args.nums), "maxFrequencyElements(%v)", tt.args.nums)
		})
	}
}
