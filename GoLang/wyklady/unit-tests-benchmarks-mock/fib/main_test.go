package main

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1}, {"2", 2, 1}, {"3", 3, 2},
		{"4", 4, 3}, {"5", 5, 5}, {"6", 6, 8},
		{"7", 7, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibNR(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1}, {"2", 2, 1}, {"3", 3, 2},
		{"4", 4, 3}, {"5", 5, 5}, {"6", 6, 8},
		{"7", 7, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibNR(tt.n); got != tt.want {
				t.Errorf("FibNR() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func BenchmarkFib(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkFibNR(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = FibNR(10)
	}
	result = r
}
