package iteration

import (
	"strings"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	// RepeatNTimes("a", b.N)
	for i := 0; i < b.N; i++ {
		Repeat("a") // Interesting implementation of the Repeat function, O(log n)
	}
}

func TestIteration(t *testing.T) {

	t.Run("Repeat the given character 5 times", func(t *testing.T) {
		got := Repeat("a")
		want := strings.Repeat("a", 5)

		assertCorrectMessage(t, got, want)
	})

	t.Run("Repeat the given character N times", func(t *testing.T) {
		got := RepeatNTimes("a", 10)
		want := strings.Repeat("a", 10)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
