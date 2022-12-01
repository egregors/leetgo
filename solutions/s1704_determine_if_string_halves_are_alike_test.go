package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_halvesAreAlike(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{
				"book",
			},
			true,
		},
		{
			"test-2",
			args{
				"textbook",
			},
			false,
		},
		{
			"test-3",
			args{
				"MerryChristmas",
			},
			false,
		},
		{
			"test-4",
			args{
				"AbCdEfGh",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, halvesAreAlike(tt.args.s), "halvesAreAlike(%v)", tt.args.s)
		})
	}
}
