# Testing

Automatic testcases make life easer for developers.
After adding a new feature, the programmer can
add new cases to the pool of the existing testcases.
And then run (corresponding) testcases to check:

* bugs introduced by the new feature
* performance of the new feature
* performance of old features
* performance of the entire system
* other things

Package testing provides support for automated testing of Go packages.
Thus we need to be familiar with it.

## Add a new feature

* add the feature code in a seperate file, e.g., fib.go
* add its testcase in a seperate file, e.g., fib_test.go
* run corresponding testcases to check important performance

## Optimize the code

* add new code in the corresponding file
* compare the performance with benchmark test

## Benchmark testing

[fib.go][fib.go] shows how to use [benchmark][benchmark] testing.

Once you update a function, run a benchmark testing.
Compare the output via benchstat, you got the changes
of performance.

This file gives two methods to calculate [Fibonacci numbers][Fibonacci]
of the Fibonacci sequence, in which each number is the sum
of the two preceding ones:
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...

As we can see, iteration method performs extremely better when N > 10 :

```golang
D:\proj\github.com\hzget\go-investigation\testing>go test -v -run='^$' -bench=Fib -count=5 -fibname=FibRecursion > recursion.txt

D:\proj\github.com\hzget\go-investigation\testing>go test -v -run='^$' -bench=Fib -count=5 -fibname=FibIteration > iteration.txt

D:\proj\github.com\hzget\go-investigation\testing>benchstat recursion.txt iteration.txt
name          old time/op  new time/op  delta
Fib/(N=2)-4   5.00ns ± 3%  2.95ns ± 3%   -40.85%  (p=0.008 n=5+5)
Fib/(N=3)-4   8.50ns ± 2%  4.58ns ± 1%   -46.04%  (p=0.008 n=5+5)
Fib/(N=4)-4   15.2ns ± 2%   5.5ns ± 1%   -63.91%  (p=0.008 n=5+5)
Fib/(N=5)-4   25.5ns ± 2%   5.3ns ± 4%   -79.04%  (p=0.008 n=5+5)
Fib/(N=10)-4   314ns ± 3%     6ns ± 1%   -97.98%  (p=0.008 n=5+5)
Fib/(N=15)-4  3.59µs ± 3%  0.01µs ± 4%   -99.74%  (p=0.008 n=5+5)
Fib/(N=20)-4  39.7µs ± 2%   0.0µs ± 6%   -99.97%  (p=0.008 n=5+5)
Fib/(N=25)-4   439µs ± 3%     0µs ± 4%  -100.00%  (p=0.008 n=5+5)
Fib/(N=30)-4  4.81ms ± 1%  0.00ms ± 1%  -100.00%  (p=0.008 n=5+5)

D:\proj\github.com\hzget\go-investigation\testing>
```

## Fuzzing

[reverse.go][reverse.go] shows how to use [fuzzing][Fuzzing] test.
The `Reverse()` function reverses a string of characters.
We can write testcases manually with our experience.
Fuzzing helps to generate unanticipated cases and report bugs,
and then we modifiy code and re-test it --- it evolves.

For example, the 1st version can reverse "abc", "hello", etc.
Fuzzing reports cases that the reversed string contains
invalid Unicode chars. And then we modify the code and get
the 2nd version.

```golang
D:\proj\github.com\hzget\go-investigation\testing>go test -v -run=Reverse -fuzz=Reverse
=== RUN   TestReverse
--- PASS: TestReverse (0.00s)
=== FUZZ  FuzzReverse
fuzz: elapsed: 1s, gathering baseline coverage: 0/197 completed
fuzz: minimizing 32-byte failing input file
fuzz: elapsed: 1s, gathering baseline coverage: 11/197 completed
--- FAIL: FuzzReverse (1.21s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:91: Reverse produced invalid UTF-8 string "\x9a\x9a\xe8"

    Failing input written to testdata\fuzz\FuzzReverse\338a30dc6777778e46b41e6b359aaa8338eacead059a6f0e640d5d1287538f77
    To re-run:
    go test -run=FuzzReverse/338a30dc6777778e46b41e6b359aaa8338eacead059a6f0e640d5d1287538f77
FAIL
exit status 1
FAIL    ltest   1.519s
```

[reverse.go]: ./reverse.go
[fib.go]: ./fib.go
[Fuzzing]: https://go.dev/doc/tutorial/fuzz
[benchmark]: https://pkg.go.dev/testing#hdr-Benchmarks
[Fibonacci]: https://en.wikipedia.org/wiki/Fibonacci_number
