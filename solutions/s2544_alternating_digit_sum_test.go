/*
	https://leetcode.com/problems/alternating-digit-sum/description/

	You are given a positive integer n. Each digit of n has a sign according to the following rules:

    The most significant digit is assigned a positive sign.
    Each other digit has an opposite sign to its adjacent digits.

	Return the sum of all digits with their corresponding sign.
*/

package solutions

import "testing"

func Test_alternateDigitSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{n: 521},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{n: 111},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{n: 886996},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternateDigitSum(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
