package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_distinctNames(t *testing.T) {
	type args struct {
		ideas []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test 1",
			args{
				[]string{"coffee", "donuts", "time", "toffee"},
			},
			6,
		},
		{
			"test 2",
			args{
				[]string{"lack", "back"},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, distinctNames(tt.args.ideas), "distinctNames(%v)", tt.args.ideas)
		})
	}
}
