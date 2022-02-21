package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{logs: []string{"d1/", "d2/", "../", "d21/", "./"}},
			2,
		},
		{
			"Test 2",
			args{logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}},
			3,
		},
		{
			"Test 3",
			args{logs: []string{"d1/", "../", "../", "../"}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minOperations(tt.args.logs), "minOperations(%v)", tt.args.logs)
		})
	}
}
