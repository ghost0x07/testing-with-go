package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	rec := Rectangle{10, 20}
	got := rec.Perimeter()
	want := 60.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rec := Rectangle{10, 20}
		checkArea(t, rec, 200)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 200)
	})

}
