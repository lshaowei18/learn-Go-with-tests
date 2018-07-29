package structs

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Rectangle Perimeter", func(t *testing.T) {

		rectangle := Rectangle{10, 10}
		got := Perimeter(rectangle)
		expected := 40.0

		if got != expected {
			t.Errorf("got %v; expected %v", got, expected)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v got %v; expected %v", test.shape, got, test.want)
			}
		})
	}

}
