package solutions

import (
	"reflect"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			[]int{1, 2},
		},
		{
			"Test 2",
			args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			[]int{1, 3},
		},
		{
			"Test 3",
			args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
