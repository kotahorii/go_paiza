package main

import (
	"fmt"
	"go/paiza/question"
)

func main() {
	win, lose := question.CountWinAndLose()
	fmt.Printf("win: %d, lose: %d", win, lose)
}
