package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_highFive(t *testing.T) {
	type args struct {
		items [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//Input: items = [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
		//Output: [[1,87],[2,88]]
		{
			name: "Test 1",
			args: args{items: [][]int{
				{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60},
				{2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76},
			}},
			want: [][]int{
				{1, 87}, {2, 88},
			},
		},
		//Input: items = [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
		//Output: [[1,100],[7,100]]
		{
			name: "Test 2",
			args: args{items: [][]int{
				{1, 100}, {7, 100}, {1, 100}, {7, 100}, {1, 100},
				{7, 100}, {1, 100}, {7, 100}, {1, 100}, {7, 100},
			}},
			want: [][]int{
				{1, 100}, {7, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, highFive(tt.args.items), "highFive(%v)", tt.args.items)
		})
	}
}
