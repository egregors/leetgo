package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name  string
		args  args
		wantS int
	}{
		//Input: n = 10, m = 3
		//Output: 19
		{"Test1", args{10, 3}, 19},
		//Input: n = 5, m = 6
		//Output: 15
		{"Test2", args{5, 6}, 15},
		//Input: n = 5, m = 1
		//Output: -15
		{"Test3", args{5, 1}, -15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantS, differenceOfSums(tt.args.n, tt.args.m), "differenceOfSums(%v, %v)", tt.args.n, tt.args.m)
		})
	}
}
