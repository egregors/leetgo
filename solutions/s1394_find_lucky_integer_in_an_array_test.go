package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{2, 2, 3, 4},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{1, 2, 2, 3, 3, 3},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				arr: []int{2, 2, 2, 3, 3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLucky(tt.args.arr), "findLucky(%v)", tt.args.arr)
		})
	}
}
