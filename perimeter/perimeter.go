package perimeter
type Circle struct{
	field float64
}
type Rectangle struct {
	width  float64
	height float64
}
func (r Rectangle) Area(){
	
}
type common interface{
	Rectangle
	Circle
}
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}
func Area(w Rectangle) float64 {
	d := Rectangle{
		width:  w.width,
		height: w.height,
	}
	return d.width * d.height
}
