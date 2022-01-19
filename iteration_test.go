package main

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("5 times call", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("3 times call", func(t *testing.T) {
		repeated := Repeat("z", 3)
		expected := "zzz"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
