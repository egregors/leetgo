package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canVisitAllRooms(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				rooms: [][]int{
					{1},
					{2},
					{3},
					{},
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				rooms: [][]int{
					{1, 3},
					{3, 0, 1},
					{2},
					{0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canVisitAllRooms(tt.args.rooms), "canVisitAllRooms(%v)", tt.args.rooms)
		})
	}
}
