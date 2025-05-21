package solutions

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
		//Output: [[1,0,1],[0,0,0],[1,0,1]]
		{
			"test1",
			args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		//Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
		//Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
		{
			"test2",
			args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if len(tt.args.matrix) != len(tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want)
				return
			}
		})
	}
}
