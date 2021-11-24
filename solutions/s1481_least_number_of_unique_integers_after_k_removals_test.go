package solutions

import "testing"

func Test_findLeastNumOfUniqueInts(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				arr: []int{5, 5, 4},
				k:   1,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				arr: []int{4, 3, 1, 1, 3, 3, 2},
				k:   3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeastNumOfUniqueInts(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findLeastNumOfUniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
