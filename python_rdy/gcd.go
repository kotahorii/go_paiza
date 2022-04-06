package pythonrdy

import (
	"log"
)

func GCD(a, b int) int {
	log.Println(a, b)
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
