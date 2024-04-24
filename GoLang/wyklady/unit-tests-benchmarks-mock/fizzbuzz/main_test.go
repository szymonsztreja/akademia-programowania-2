package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "3", n: 3, want: "Fuzz"},
		{name: "5", n: 5, want: "Buzz"},
		{name: "15", n: 15, want: "FizzBuzz"},
		{name: "-3", n: -3, want: "Fizz"},
		{name: "0", n: 0, want: "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.n); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
