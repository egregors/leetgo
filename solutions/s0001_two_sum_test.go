package solutions

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case", args{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		}, []int{0, 1}},
		{"Case", args{
			nums:   []int{3, 2, 4},
			target: 6,
		}, []int{1, 2}},
		{"Case", args{
			nums:   []int{3, 3},
			target: 6,
		}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
