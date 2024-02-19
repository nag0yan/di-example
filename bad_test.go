package main

import "testing"

func TestHelloFromBadX(t *testing.T) {
	// the result is "hello", and this is never changed (unflexible)
	got := helloFromBadX()
	want := "hello world"
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
