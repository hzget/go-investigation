package ltest

import (
	"testing"
)

func testFibN(t *testing.T, fib func(int) int, n int, want int) {
	if got := fib(n); got != want {
		t.Errorf("Fib(%d) want %d, but got %d", n, want, got)
	}
}

func testFib(t *testing.T, fib func(int) int) {
	cases := []struct {
		in   int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}}

	for _, v := range cases {
		if got := fib(v.in); got != v.want {
			t.Fatalf("Fib(%d) want %d, but got %d", v.in, v.want, got)
		}
	}
}

func TestFibRecursion(t *testing.T) {
	testFib(t, FibRecursion)
}

func TestFibIteration(t *testing.T) {
	testFib(t, FibIteration)
}

func benchmarkFib(b *testing.B, fib func(int) int, n int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fib(n)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	b.Run("Iteration(10)", benchmarkFib(b, FibIteration, 10))
	b.Run("Recursion(10)", benchmarkFib(b, FibRecursion, 10))
	b.Run("Iteration(30)", benchmarkFib(b, FibIteration, 30))
	b.Run("Recursion(30)", benchmarkFib(b, FibRecursion, 30))
}

func BenchmarkFibIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIteration(30)
	}
}

func BenchmarkFibRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursion(30)
	}
}
