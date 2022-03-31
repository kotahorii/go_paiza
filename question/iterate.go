package question

import (
	"fmt"
	"math"
)

func DisplaySPAM10Times() []string {
	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, "SPAM")
	}
	return result
}

func DisplayThreeTimes() {
	for i := 1; i < 10; i++ {
		fmt.Printf("3 * %d = %d\n", i, 3*i)
	}
}

func DisplayTwoToThePower() {
	for i := 1; i < 9; i++ {
		fmt.Printf("2 to the power of %d = %d\n", i, int(math.Pow(2, float64(i))))
	}
}
