package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_poorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"t0",
			args{buckets: 1000, minutesToDie: 15, minutesToTest: 60},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest), "poorPigs(%v, %v, %v)", tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest)
		})
	}
}
