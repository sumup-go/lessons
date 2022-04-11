package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	out := Reverse("👪")
	expected := "👪"
	if out != expected {
		t.Fatalf("expected %s, got %s\n", expected, out)
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345", "👪"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, s string) {
		rev := Reverse(s)
		doubleRev := Reverse(rev)
		if s != doubleRev {
			t.Errorf("Before: %q, after: %q", s, doubleRev)
		}
		if utf8.ValidString(s) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
