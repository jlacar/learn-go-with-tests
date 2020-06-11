package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("empty name defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty language defaults to English", func(t *testing.T) {
		got := Hello("Tom", "")
		want := "Hello, Tom!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("unknown language defaults to English", func(t *testing.T) {
		got := Hello("Witcher", "High Valerian")
		want := "Hello, Witcher!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Hello' in English", func(t *testing.T) {
		got := Hello("Junilu", "English")
		want := "Hello, Junilu!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Hola' in Spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Bonjour' in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Shalom' in Hebrew", func(t *testing.T) {
		got := Hello("Yitzakh", "Hebrew")
		want := "Shalom, Yitzakh!"
		assertCorrectMessage(t, got, want)
	})
}
