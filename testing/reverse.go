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

// reverse character orders in a string
func Reverse(s string) (string, error) {
	// return Reverse1(s), nil
	// return Reverse2(s), nil
	return Reverse3(s)
}
