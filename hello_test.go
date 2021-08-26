package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Faiz", func(t *testing.T) {
		got := hello("Faiz")
		want := "Hello, Faiz!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World!', when an empty string is provided", func(t *testing.T) {
		got := hello("")
		want := "Hello, World!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
