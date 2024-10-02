package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello World", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, origin string) {
		rev, err1 := Reverse(origin)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		// rev := Reverse(origin)
		// doubleRev := Reverse(rev)

		t.Logf("Number of runes: origin=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(origin), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if origin != doubleRev {
			t.Errorf("Before: %q, after: %q", origin, doubleRev)
		}

		if utf8.ValidString(origin) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}

	})
}
