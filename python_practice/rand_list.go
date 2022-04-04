package pythonpractice

import "math/rand"

func RandList(leng, lower, upper int) (result []int) {
	for len(result) < leng {
		value := rand.Intn(100)

		if value > lower && value < upper {
			result = append(result, value)
		}
	}
	return
}
