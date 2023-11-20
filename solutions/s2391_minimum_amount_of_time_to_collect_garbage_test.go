package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_garbageCollection(t *testing.T) {
	type args struct {
		garbage []string
		travel  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 0",
			args{
				[]string{"G", "P", "GP", "GG"},
				[]int{2, 4, 3},
			},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, garbageCollection(tt.args.garbage, tt.args.travel), "garbageCollection(%v, %v)", tt.args.garbage, tt.args.travel)
		})
	}
}
