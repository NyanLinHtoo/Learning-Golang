package main

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Errorf("Expected True but got false")
	}
}

func FuzzEqaual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b) // fuzz target is support for one target
	})
}
