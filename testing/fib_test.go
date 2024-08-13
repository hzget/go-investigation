package ltest

import (
	"testing"
)

var seq = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

func testFib(t *testing.T, f func(int) int) func(*testing.T) {
	return func(t *testing.T) {
		for n, v := range seq {
			if got := f(n); got != v {
				t.Fatalf("fib(%d) expects %d, but got %d", n, v, got)
			}
		}
	}
}

func TestFib(t *testing.T) {
	t.Run("Recursion", testFib(t, FibRecursion))
	t.Run("Iteration", testFib(t, FibIteration))
	t.Run("Fib", testFib(t, Fib))
}
