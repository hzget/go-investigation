/*
This file shows how to use benchmark testing.

Once you update a function, run a benchmark testing.
Compare the output via benchstat, you got the changes of performance.
For example:

# go test -v -run='^$' -bench=Fib -count=5 -fibname=FibRecursion > recursion.txt
# go test -v -run='^$' -bench=Fib -count=5 -fibname=FibIteration > iteration.txt
# benchstat recursion.txt iteration.txt

name          old time/op  new time/op  delta
Fib/(N=2)-4   5.00ns ± 3%  2.95ns ± 3%   -40.85%  (p=0.008 n=5+5)
Fib/(N=3)-4   8.50ns ± 2%  4.58ns ± 1%   -46.04%  (p=0.008 n=5+5)
Fib/(N=4)-4   15.2ns ± 2%   5.5ns ± 1%   -63.91%  (p=0.008 n=5+5)
Fib/(N=5)-4   25.5ns ± 2%   5.3ns ± 4%   -79.04%  (p=0.008 n=5+5)
Fib/(N=10)-4   314ns ± 3%     6ns ± 1%   -97.98%  (p=0.008 n=5+5)
Fib/(N=15)-4  3.59μs ± 3%  0.01μs ± 4%   -99.74%  (p=0.008 n=5+5)
Fib/(N=20)-4  39.7μs ± 2%   0.0μs ± 6%   -99.97%  (p=0.008 n=5+5)
Fib/(N=25)-4   439μs ± 3%     0μs ± 4%  -100.00%  (p=0.008 n=5+5)
Fib/(N=30)-4  4.81ms ± 1%  0.00ms ± 1%  -100.00%  (p=0.008 n=5+5)
*/
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
