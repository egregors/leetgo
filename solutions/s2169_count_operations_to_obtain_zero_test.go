package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				num1: 2,
				num2: 3,
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				num1: 10,
				num2: 10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countOperations(tt.args.num1, tt.args.num2), "countOperations(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
