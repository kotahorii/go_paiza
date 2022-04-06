package pythonrdy

import (
	"fmt"
	"math"
)

func PrimeNumber(max int) (result []int) {
	var list []int
	for i := 2; i < max+1; i++ {
		list = append(list, i)
	}

	for i := 2; i < int(math.Pow(float64(max), 0.5))+1; i++ {
		for j, n := range list {
			if n == i || n%i != 0 {
				continue
			}
			list[j] = 0
		}
	}
	for _, n := range list {
		if n != 0 {
			result = append(result, n)
		}
	}
	return
}

func NTimes(n float64) string {
	num := int(math.Pow(n, math.Pow(n, n)))
	result := fmt.Sprint(num)[1]
	return string(result)
}
