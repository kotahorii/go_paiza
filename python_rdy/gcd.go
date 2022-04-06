package pythonrdy

func GCP(a, b int) int {
	if b == 0 {
		return a
	}
	return GCP(b, a%b)
}
