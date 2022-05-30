package solutions

import (
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "Test 3",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		},
		{
			name: "Test 4",
			args: args{
				dividend: -2147483648,
				divisor:  1,
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
