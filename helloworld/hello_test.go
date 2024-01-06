package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people ", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("World", "English")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mariano", "Spanish")
		want := "Hola Mariano"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour Pierre"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Francis", "Italian")
		want := "Ciao Francis"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
