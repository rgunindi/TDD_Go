package iteration

import "strings"

func Iterate(v string, c int) string {
return strings.Repeat(v, c)
}
func BadIterate(v string, c int) string {
	var repeat string
	for i := 0; i < c; i++ {
		repeat+=v
	}
	return repeat
}
