package main

import "go/paiza/question"

func main() {
	cat := question.NewCat("takashi", 43)
	question.PrintAnimalName(cat)
}
