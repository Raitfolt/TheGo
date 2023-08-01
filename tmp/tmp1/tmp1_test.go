package main

import (
	"testing"
)

func Test_gcd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10 & 5",
			args: args{x: 10, y: 5},
			want: 5,
		},
		{
			name: "16 & 88",
			args: args{x: 16, y: 88},
			want: 8,
		},
		{
			name: "666 & 222",
			args: args{x: 666, y: 222},
			want: 222,
		},
		{
			name: "221 & 323",
			args: args{x: 221, y: 323},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fibonacci`s number 1",
			n:    1,
			want: 0,
		},
		{
			name: "Fibonacci`s number 2",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci`s number 3",
			n:    3,
			want: 1,
		},
		{
			name: "Fibonacci`s number 4",
			n:    4,
			want: 2,
		},
		{
			name: "Fibonacci`s number 5",
			n:    5,
			want: 3,
		},
		{
			name: "Fibonacci`s number 10",
			n:    10,
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
