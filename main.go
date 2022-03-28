package main

import "fmt"

func main() {
	divideNumber(13)
}

func printInt(x, y int) {
	fmt.Printf("x=%d, y=%d", x, y)
}

func divideNumber(num int) {
	for i := 1; i < num-1; i++ {
		fmt.Printf("remainder %d, quotient %d\n", num%i, num/i)
	}
}
