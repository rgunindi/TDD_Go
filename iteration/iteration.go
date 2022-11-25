package iteration


func Iterate(v string,c int )string{
	var repeated string
	for c>0{
		repeated+=v
		c--
	}
	return repeated;
}
