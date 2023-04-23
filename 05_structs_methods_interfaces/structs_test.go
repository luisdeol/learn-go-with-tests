package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if expected != got {
		t.Errorf("expected %.2f but got %.2f instead", expected, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, expected: 100},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 5}, expected: 25},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if tt.expected != got {
			t.Errorf("expected %2.f but got %.2f instead", tt.expected, got)
		}
	}
}
