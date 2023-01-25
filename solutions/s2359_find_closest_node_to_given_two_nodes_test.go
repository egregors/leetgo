package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_closestMeetingNode(t *testing.T) {
	type args struct {
		edges []int
		node1 int
		node2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{2, 2, 3, -1},
				0,
				1,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, closestMeetingNode(tt.args.edges, tt.args.node1, tt.args.node2), "closestMeetingNode(%v, %v, %v)", tt.args.edges, tt.args.node1, tt.args.node2)
		})
	}
}
