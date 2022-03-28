package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var x, y int
	if sc.Scan() {
		if num, err := strconv.Atoi(sc.Text()); err != nil {
			log.Fatal("input invalid value")
		} else {
			x = num
		}
	}

	if sc.Scan() {
		if num, err := strconv.Atoi(sc.Text()); err != nil {
			log.Fatal("input invalid value")
		} else {
			y = num
		}
	}

	PrintAddAndSub(x, y)
}

func PrintAddAndSub(x, y int) {
	fmt.Printf("sum is %d\n", x+y)
	fmt.Printf("sub is %d", x-y)
}
