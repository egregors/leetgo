package solutions

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			4,
		},
		{
			"Test",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			-1,
		},
		{
			"Test",
			args{
				nums:   []int{5},
				target: 5,
			},
			0,
		},
		{
			"Test",
			args{
				nums:   []int{5},
				target: -5,
			},
			-1,
		},
		{
			"Test",
			args{
				nums:   []int{2, 5},
				target: 5,
			},
			1,
		},
		{
			"Test",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 13,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
