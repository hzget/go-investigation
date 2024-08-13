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
	"testing"
	"unicode/utf8"
)

var cases = []struct {
	in   string
	want string
}{
	{"hello", "olleh"},
	{"haha!", "!ahah"},
	{"", ""},
	{" ", " "},
	{"@^", "^@"},
	{"'`\"", "\"`'"},
}
var seeds = []string{"hello", "haha!", "`\"@", "", " "}

func TestReverse(t *testing.T) {
	for _, v := range cases {
		got, err := Reverse(v.in)
		if err != nil {
			t.Error(err)
		}
		if got != v.want {
			t.Errorf("want %s, got %s", v.want, got)
		}
	}
}

func FuzzReverse(f *testing.F) {
	for _, v := range seeds {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, origin string) {
		rev, err := Reverse(origin)
		if err != nil {
			t.Skip("skip this case: ", err)
		}
		revrev, _ := Reverse(rev)
		if revrev != origin {
			t.Errorf("origin %q, rev %q, revrev %q", origin, rev, revrev)
		}
		if utf8.ValidString(origin) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
