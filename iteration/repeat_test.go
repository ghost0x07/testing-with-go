package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q, got %q", expected, repeated)
		}
	})
	t.Run("bb 4 times", func(t *testing.T) {
		repeated := Repeat("bb", 4)
		expected := "bbbbbbbb"

		if repeated != expected {
			t.Errorf("expected %q, got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
