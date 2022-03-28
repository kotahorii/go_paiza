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

func PrintAddAndSub(x, y int) {
	fmt.Printf("sum is %d\n", x+y)
	fmt.Printf("sub is %d", x-y)
}

func PrintMeanValue(x, y int) {
	fmt.Println((x + y) / 2)
}
