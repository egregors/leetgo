package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_guessNumber(t *testing.T) {
	type args struct {
		n     int
		guess func(int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				n:     10,
				guess: makeGuess(6),
			},
			6,
		},
		{
			"test 2",
			args{
				n:     1,
				guess: makeGuess(1),
			},
			1,
		},
		{
			"test 3",
			args{
				n:     2,
				guess: makeGuess(1),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, guessNumber(tt.args.n, tt.args.guess), "guessNumber(%v, %v)", tt.args.n, tt.args.guess)
		})
	}
}
