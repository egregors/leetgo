package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: num = "6777133339"
		//Output: "777"
		//Input: num = "2300019"
		//Output: "000"
		//Input: num = "42352338"
		//Output: ""
		{
			name: "Test 1",
			args: args{num: "6777133339"},
			want: "777",
		},
		{
			name: "Test 2",
			args: args{num: "2300019"},
			want: "000",
		},
		{
			name: "Test 3",
			args: args{num: "42352338"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestGoodInteger(tt.args.num), "largestGoodInteger(%v)", tt.args.num)
		})
	}
}
