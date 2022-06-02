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

