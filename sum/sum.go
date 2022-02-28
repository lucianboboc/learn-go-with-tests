package sum

func Sum(a []int) int {
	var sum = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func SumAll(s ...[]int) []int {
	var result []int
	for _, sl := range s {
		result = append(result, Sum(sl))
	}
	return result
}

func SumAllTail(s ...[]int) []int {
	var result []int
	for _, sl := range s {
		if len(sl) == 0 {
			result = append(result, 0)
		} else {
			tail := sl[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}
