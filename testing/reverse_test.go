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
