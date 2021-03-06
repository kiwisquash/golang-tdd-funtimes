package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ki")
	want := "Hello, Ki!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}