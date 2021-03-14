package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"hello": "world!"}

	t.Run("in the dict", func(t *testing.T) {
		got, err := dictionary.Search("hello")
		want := "world!"

		assertStrings(t, got, want, dictionary)
		assertError(t, err, nil, dictionary)
	})

	t.Run("not in the dict", func(t *testing.T) {
		got, err := dictionary.Search("bob")
		want := ErrNotFound

		assertStrings(t, got, "", dictionary)
		assertError(t, err, want, dictionary)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new word", func (t *testing.T) {
		dictionary := Dictionary{}

		word := "hello"
		definition := "world!"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil, dictionary)
		assertDefinition(t, word, definition, dictionary)
	})

	t.Run("add existing word", func (t *testing.T) {

		word := "hello"
		definition := "world!"

		dictionary := Dictionary{word:definition}

		err := dictionary.Add(word, "fish")

		assertError(t, err, ErrWordExists, dictionary)
		assertDefinition(t, word, definition, dictionary)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("word already exists", func(t *testing.T) {
		word := "hello"
		definition := "world!"

		dictionary := Dictionary{word:definition}
		
		newDefinition := "Steve"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil, dictionary)
		assertDefinition(t, word, newDefinition, dictionary)
	})

	t.Run("word doesn't exist", func(t *testing.T) {

		dictionary := Dictionary{}

		word := "hello"
		definition := "world!"

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist, dictionary)
	})
}

func TestDelete (t *testing.T) {
	dictionary := Dictionary{"hello": "world!"}

	dictionary.Delete("hello")

	_, err := dictionary.Search("hello")
	assertError(t, err, ErrNotFound, dictionary)
}

func assertDefinition(t testing.TB, word, definition string, dictionary Dictionary) {

	t.Helper()
	
	want := definition
	got, error := dictionary.Search(word)

	if error != nil {
		t.Fatal("Shouldn't have an error.")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	} 

}

func assertStrings(t testing.TB, got, want string, dictionary Dictionary) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, dictionary)
	}

}

func assertError(t testing.TB, got, want error, dictionary Dictionary) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, dictionary)
	}
}