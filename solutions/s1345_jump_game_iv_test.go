package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minJumps(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minJumps(tt.args.arr), "minJumps(%v)", tt.args.arr)
		})
	}
}
