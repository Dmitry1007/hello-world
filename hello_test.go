package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	want := "Howzit?"
	if got := Greet(); got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
