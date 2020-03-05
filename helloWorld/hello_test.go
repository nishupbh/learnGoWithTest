package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Nishant", "")
		want := "Hello, Nishant"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, World' for spanish when an empty string is being supplied", func(t *testing.T) {
		got := hello("", "spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, World' when an empty string is being supplied", func(t *testing.T) {
		got := hello("Nishant", "french")
		want := "Bonjour, Nishant"
		assertCorrectMessage(t, got, want)
	})
}
