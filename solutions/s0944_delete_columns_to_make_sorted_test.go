package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]string{"cba", "daf", "ghi"},
			},
			1,
		},
		{
			"case 2",
			args{
				[]string{"a", "b"},
			},
			0,
		},
		{
			"case 3",
			args{
				[]string{"zyx", "wvu", "tsr"},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDeletionSize(tt.args.strs), "minDeletionSize(%v)", tt.args.strs)
		})
	}
}
