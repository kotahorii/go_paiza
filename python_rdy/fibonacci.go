package pythonrdy

func Fib(n int) (result []int) {
	a, b := 0, 1
	if n == 1 {
		return []int{a}
	}
	if n == 2 {
		return []int{a, b}
	}

	result = []int{a, b}
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
		result = append(result, b)
	}
	return
}
