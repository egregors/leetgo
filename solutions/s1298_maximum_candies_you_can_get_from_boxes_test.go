package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxCandies(t *testing.T) {
	type args struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Input: status = [1,0,1,0], candies = [7,5,4,100], keys = [[],[],[1],[]], containedBoxes = [[1,2],[3],[],[]], initialBoxes = [0]
		//Output: 16
		{
			"Test 1",
			args{
				status:         []int{1, 0, 1, 0},
				candies:        []int{7, 5, 4, 100},
				keys:           [][]int{{}, {}, {1}, {}},
				containedBoxes: [][]int{{1, 2}, {3}, {}, {}},
				initialBoxes:   []int{0},
			},
			16,
		},
		//Input: status = [1,0,0,0,0,0], candies = [1,1,1,1,1,1], keys = [[1,2,3,4,5],[],[],[],[],[]], containedBoxes = [[1,2,3,4,5],[],[],[],[],[]], initialBoxes = [0]
		//Output: 6
		{
			"Test 2",
			args{
				status:         []int{1, 0, 0, 0, 0, 0},
				candies:        []int{1, 1, 1, 1, 1, 1},
				keys:           [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
				containedBoxes: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
				initialBoxes:   []int{0},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxCandies(tt.args.status, tt.args.candies, tt.args.keys, tt.args.containedBoxes, tt.args.initialBoxes), "maxCandies(%v, %v, %v, %v, %v)", tt.args.status, tt.args.candies, tt.args.keys, tt.args.containedBoxes, tt.args.initialBoxes)
		})
	}
}
