package synccounter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing Counter 3 times should leave it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, &counter, 3)

	})

	t.Run("it runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup

		for i := 0; i < wantedCount; i++ {
			wg.Add(1)
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCount(t, &counter, wantedCount)
	})
}

func assertCount(t testing.TB, counter *Counter, expected int) {
	if counter.Value() != expected {
		t.Errorf("Counts: got %d, want %d", counter.Value(), expected)
	}
}
