package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is test dictionary"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is test dictionary"
		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal(" expected to get an error")
		}
		assertString(t, err.Error(), want)
	})
}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		definition := "this is just a test"

		if err != nil {
			t.Fatal("Should find the added word", err)
		}
		assertString(t, err, nil)
		assertDefinition(t, dictionary, "test", definition)

	})

	t.Run("Existing word", func(t *testing.T) {
		definition := "this is just a test"
		dictionary := Dictionary{"test": definition}
		err := dictionary.Add("test", "new test")
		assertString(t, err, ErrWordExists)
		assertDefinition(t, dictionary, "test", definition)
	})

}

func assertString(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDefinition(t *testing.T, d Dictionary, word string, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("Should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
