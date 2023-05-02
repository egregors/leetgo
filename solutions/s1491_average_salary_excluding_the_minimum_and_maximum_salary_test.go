package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"case 1",
			args{[]int{4000, 3000, 1000, 2000}},
			2500.00000,
		},
		{
			"case 2",
			args{[]int{1000, 2000, 3000}},
			2000.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, average(tt.args.salary), "average(%v)", tt.args.salary)
		})
	}
}
