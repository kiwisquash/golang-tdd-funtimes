package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Kiwi")

	got := buffer.String()
	want := "Hello, Kiwi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

