package pythonrdy

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

func Sum(n int) int {
	return sum(10)
}

