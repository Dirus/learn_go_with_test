package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.00, 5.00}
	got := rectangle.Perimeter()
	want := 14.00
	if got != want {
		t.Errorf("Want %f but got %f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.00},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("Want %f but got %f", tt.hasArea, got)
			}
		})
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("Want %f but got %f", want, got)
// 		}
// 	}
// 	t.Run("CheckArea for rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{2.00, 5.00}
// 		want := 10.00
// 		checkArea(t, rectangle, want)
// 	})
// 	t.Run("CheckArea for Circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }
