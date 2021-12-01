package solutions

import "testing"

func Test_maxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				},
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				grid: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 1, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
