package question

import "fmt"

func PrintInt(x, y int) {
	fmt.Printf("x=%d, y=%d", x, y)
}

func DivideNumber(num int) {
	for i := 1; i < num-1; i++ {
		fmt.Printf("remainder %d, quotient %d\n", num%i, num/i)
	}
}
