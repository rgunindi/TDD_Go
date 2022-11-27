package perimeter

import "math"

type Circle struct {
	Radius float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	d := Rectangle{
		width:  r.width,
		height: r.height,
	}
	return d.width * d.height
}
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2.0) * math.Pi
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}
func Area(r Rectangle) float64 {
	return 0
}
