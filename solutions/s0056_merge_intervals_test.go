package solutions

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Test 2",
			args: args{
				intervals: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "Test 3",
			args: args{
				intervals: [][]int{
					{1, 4},
					{0, 4},
				},
			},
			want: [][]int{
				{0, 4},
			},
		},
		{
			name: "Test 4",
			args: args{
				intervals: [][]int{
					{1, 4},
					{0, 2},
					{3, 5},
				},
			},
			want: [][]int{
				{0, 5},
			},
		},
		{
			name: "Test 5",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 3},
				},
			},
			want: [][]int{
				{1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
