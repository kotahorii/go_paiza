package main

import (
	"fmt"
	pythonrdy "go/paiza/python_rdy"
)

func main() {
	fmt.Println(pythonrdy.CountPython("./text.txt", "python"))
}
