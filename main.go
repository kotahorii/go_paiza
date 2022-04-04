package main

import (
	"fmt"
	pythonpractice "go/paiza/python_practice"
	"strings"
)

func main() {
	fmt.Println(pythonpractice.MapConcat(func(s string) string { return strings.Repeat(s, 3) }, "abc", "-"))
}
