package ltest

import (
	"flag"
	"os"
	"testing"
)

func benchmarkFib(b *testing.B, fib func(int) int, n int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fib(n)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	b.Run("(N=2)", benchmarkFib(b, fib, 2))
	b.Run("(N=3)", benchmarkFib(b, fib, 3))
	b.Run("(N=4)", benchmarkFib(b, fib, 4))
	b.Run("(N=5)", benchmarkFib(b, fib, 5))
	b.Run("(N=10)", benchmarkFib(b, fib, 10))
	b.Run("(N=15)", benchmarkFib(b, fib, 15))
	b.Run("(N=20)", benchmarkFib(b, fib, 20))
	b.Run("(N=25)", benchmarkFib(b, fib, 25))
	b.Run("(N=30)", benchmarkFib(b, fib, 30))
}

var fib func(int) int

func TestMain(m *testing.M) {
	// choose the fib function from command line
	var fibname string
	flag.StringVar(&fibname, "fibname", "FibIteration",
		"mandatary. choose a fib func: FibIteration, FibRecursion")
	flag.Parse()

	fibs := make(map[string](func(int) int))
	fibs["FibRecursion"] = FibRecursion
	fibs["FibIteration"] = FibIteration
	var ok bool
	fib, ok = fibs[fibname]
	if !ok {
		flag.PrintDefaults()
		return
	}

	os.Exit(m.Run())
}
