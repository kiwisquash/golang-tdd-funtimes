package main

import "testing"

func TestHello(t *testing.T) {

	// What is this testing.TB?
	// For helper functions, just use it since it works with both T and B?
	assertCorrectMessage := func (got string, want string, t testing.TB )  {
		t.Helper() // tells testing suite that this is a helper function.

		if got != want {
			t.Errorf("got %q want %q", got, want) // without t.Helper(), test suite reports this line
		}
	}

	t.Run("say hello to a person", func(t *testing.T) {
		got := Hello("Ki", "English")
		want := "Hello, Ki!"

		assertCorrectMessage(got, want, t)

	})

	t.Run("when '' is supplied say hello to world", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"

		assertCorrectMessage(got, want, t)
	})

	t.Run("How about some Korean?", func(t *testing.T) {
		got := Hello("붕어", "Korean")
		want := "안녕, 붕어!"

		assertCorrectMessage(got, want, t)
	})

	t.Run("How about some Spanish?", func(t *testing.T) {
		got := Hello("Clara", "Spanish")
		want := "¡Hola, Clara!"

		assertCorrectMessage(got, want, t)
	})

	t.Run("How about some French?", func(t *testing.T) {
		got := Hello("Clare", "French")
		want := "Bonjour, Clare!"

		assertCorrectMessage(got, want, t)
	})
}