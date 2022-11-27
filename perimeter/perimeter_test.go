package perimeter

import "testing"

// run
// go test -run TestArea/Rectangle
func TestPerimeter(t *testing.T) {
	got := Perimeter(struct {
		width  float64
		height float64
	}{
		width:  10.0,
		height: 10.0,
	})
	hasArea := 40.0
	if got != hasArea {
		t.Errorf("got %.2f hasArea %.2f", got, hasArea)
	}
}

//anonymous struct for rectangle
// struct{
// 	width  float64
// 	height float64
// }{
// 	width:  12.0,
// 	height: 6.0,
// },

func TestArea(t *testing.T) {
	//Table driven tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 122, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		//we used tt.name for t.Run spec name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

// {
// 	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != hasArea {
// 			t.Errorf("got %g hasArea %g", got, hasArea)
// 		}
// 	}
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{width: 12.0, height: 6.0}
// 		//Function implementing
// 		// got := Area(rectangle)
// 		//method implementing
// 		hasArea := 72.0
// 		checkArea(t, rectangle, hasArea)
// 	})
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		hasArea := 314.1592653589793
// 		checkArea(t, circle, hasArea)
// 	})

// }
