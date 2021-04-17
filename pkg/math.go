package pkg

func Sum(list ...int) int {
	var result int
	for _, v := range list {
		result += v
	}
	return result
}
