package structs

import (
	"math"
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
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{
			desc:  "Circle: Radius 10",
			shape: Circle{10},
			want:  math.Pi * 10 * 10,
		},
		{
			desc:  "Rectangle: 10X20",
			shape: Rectangle{10, 20},
			want:  10 * 20,
		},
		{
			desc:  "Triangle",
			shape: Triangle{12, 6},
			want:  36,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.shape.Area()
			if got != tC.want {
				t.Errorf("%#v got %g, want %g", tC.shape, got, tC.want)
			}
		})
	}
}
