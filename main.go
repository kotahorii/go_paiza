package main

import (
	"fmt"
	pythonpractice "go/paiza/python_practice"
)

func main() {
	fmt.Println(pythonpractice.MyFilter(pythonpractice.IsEven, []int{1, 2, 3, 4, 5}))
}
