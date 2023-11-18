package leetcode

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "42"},
			want: 42,
		}, {
			name: "Example 2",
			args: args{s: "   -42"},
			want: -42,
		}, {
			name: "Example 3",
			args: args{s: "4193 with words"},
			want: 4193,
		}, {
			name: "Example 4",
			args: args{s: "words and 987"},
			want: 0,
		}, {
			name: "Example 5",
			args: args{s: "-2147483649 yahooo"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
