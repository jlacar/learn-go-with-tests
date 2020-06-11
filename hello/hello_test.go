package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Junilu")
	want := "Hello, Junilu!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
