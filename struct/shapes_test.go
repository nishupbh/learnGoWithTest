package structs

import (
	"testing"
)

func TestShapes(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %0.2f and got %0.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.area()
		if got != want {
			t.Errorf("want %g and got %g", want, got)
		}
	}

	t.Run("Rectange", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("Rectange", func(t *testing.T) {

		circle := Circle{10.0}

		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func TestAreaT(t *testing.T) {
	areaTest := []struct {
		s    Shape
		want float64
	}{
		{Rectangle{12, 5}, 60.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{10, 5}, 25.0},
	}

	for _, tt := range areaTest {
		if tt.s.area() != tt.want {
			t.Errorf("want %g and got %g", tt.want, tt.s.area())
		}
	}
}
