package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 1, 2, 4},
			},
			want: []int{4, 2, 1, 3},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 3, 2, 4},
			},
			want: []int{4, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortArrayByParity(tt.args.nums), "sortArrayByParity(%v)", tt.args.nums)
		})
	}
}
