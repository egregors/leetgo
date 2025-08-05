package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Input: fruits = [4,2,5], baskets = [3,5,4]
		//
		//Output: 1
		//Input: fruits = [3,6,1], baskets = [6,4,7]
		//
		//Output: 0
		{
			name: "Example 1",
			args: args{
				fruits:  []int{4, 2, 5},
				baskets: []int{3, 5, 4},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				fruits:  []int{3, 6, 1},
				baskets: []int{6, 4, 7},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numOfUnplacedFruits(tt.args.fruits, tt.args.baskets), "numOfUnplacedFruits(%v, %v)", tt.args.fruits, tt.args.baskets)
		})
	}
}
