package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test 1",
			args{3},
			[]string{"1", "2", "Fizz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fizzBuzz(tt.args.n), "fizzBuzz(%v)", tt.args.n)
		})
	}
}
