package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"hello": "world!"}

	t.Run("in the dict", func(t *testing.T) {

		got, err := dictionary.Search("hello")
		want := "world!"

		assertStrings(t, got, want, dictionary)

		if err != nil {
			t.Error("Shouldn't have an error.")
		}

	})

	t.Run("not in the dict", func(t *testing.T) {

		_, err := dictionary.Search("bob")
		want := "bob not found"

		if err == nil {
			t.Fatal("Was expecting an error")  // fatal stops execution
		}

		assertStrings(t, err.Error(), want, dictionary)


	})


}

func assertStrings(t testing.TB, got, want string, dictionary map[string]string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, dictionary)
	}

}