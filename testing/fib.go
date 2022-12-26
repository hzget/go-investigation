/*
This file shows how to use benchmark testing.

Once you update a function, run a benchmark testing.
Compare the output via benchstat, you got the changes
of performance.
*/
package ltest

func FibRecursion(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursion(n-1) + FibRecursion(n-2)
}

func FibIteration(n int) int {
	if n < 2 {
		return n
	}

	x, y := 0, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

func Fib(n int) int {
	// return FibRecursion(n)
	return FibIteration(n)
}
