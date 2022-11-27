package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(struct {
		width  float64
		height float64
	}{
		width:  10.0,
		height: 10.0,
	})
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{width: 12.0, height: 6.0}
		//Function implementing
		// got := Area(rectangle)

		//method implementing
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {

		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
