package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFinalValue(t *testing.T) {
	type args struct {
		nums     []int
		original int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0",
			args: args{
				nums:     []int{5, 3, 6, 1, 12},
				original: 3,
			},
			want: 24,
		},
		{
			name: "Test 1",
			args: args{
				nums:     []int{2, 7, 9},
				original: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findFinalValue(tt.args.nums, tt.args.original), "findFinalValue(%v, %v)", tt.args.nums, tt.args.original)
		})
	}
}
