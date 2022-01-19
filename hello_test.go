package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Safak", "French")
		want := "Bonjour, Safak"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Safak", "Turkish")
		want := "Selam, Safak"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Safak", "English")
		want := "Hello, Safak"
		assertCorrectMessage(t, got, want)
	})
}
