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
[Fuzzing]: https://go.dev/doc/tutorial/fuzz
