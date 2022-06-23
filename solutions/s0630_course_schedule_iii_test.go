package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{courses: [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, scheduleCourse(tt.args.courses), "scheduleCourse(%v)", tt.args.courses)
		})
	}
}
