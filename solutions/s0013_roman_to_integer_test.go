package solutions

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{s: "III"}, 3},
		{"4", args{s: "IV"}, 4},
		{"9", args{s: "IX"}, 9},
		{"58", args{s: "LVIII"}, 58},
		{"1994", args{s: "MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
