package sum

func Sum(v []int) int {
	x := 0
	for _, val := range v {
		x += val
	}
	return x
}

func SumAll(p ...[]int) []int {
	var x []int
	for _, val := range p {
		x = append(x, Sum(val))
	}
	return x
}

func SumAllTails(p ...[]int) []int {
	var sumAllTails []int
	for _, val := range p {
		if ok := len(val) == 0; ok {
			sumAllTails = append(sumAllTails, 0)
			continue
		}
		sumAllTails = append(sumAllTails, Sum(val[1:]))
	}
	return sumAllTails
}
