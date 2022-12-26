/*
This file intends to show how to use fuzzing testing.

Fuzzing is a testing technique where a function is called with
randomly generated inputs to find bugs not anticipated
by unit tests.

Links: https://go.dev/doc/tutorial/fuzz

We try to reverse a string of characters as an example. And we
expect that the result is a string of valid Unicode characters.

Fuzzing helps to generate unanticipated cases and report bugs,
and then we modifiy code and re-test it --- it evolves.

Reverse1:
 can reverse "abc", "hello", etc. Fuzzing reports cases that
 the reversed string contains invalid Unicode chars.
Reverse2:
 can reverse "你好" correctly. Fuzzing reports cases that
 the original string contains invalid Unicode chars.
Reverse3: avoid reversing invalid utf8 characters
Reverse: wrap specific version and run a fuzzing test
*/

package ltest

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse1(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// update: can reverse "你好" correctly
func Reverse2(s string) string {
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

// update: avoid reversing invalid utf8 characters
func Reverse3(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", errors.New(fmt.Sprintf("%q is invalid utf8", s))
	}

	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev), nil
}

func Reverse(s string) (string, error) {
	// return Reverse1(s), nil
	// return Reverse2(s), nil
	return Reverse3(s)
}
