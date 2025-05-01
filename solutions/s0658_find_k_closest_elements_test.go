package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test Case 2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findClosestElements(tt.args.arr, tt.args.k, tt.args.x), "findClosestElements(%v, %v, %v)", tt.args.arr, tt.args.k, tt.args.x)
		})
	}
}
