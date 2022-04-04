package main

import (
	"fmt"
	"go/paiza/question"
	"strings"
)

func main() {
	cat1 := question.NewCat("cat1", 43)
	cat2 := question.NewCat("cat2", 32)
	dog1 := question.NewDog2("dog1", 33)
	dog2 := question.NewDog2("dog2", 32)

	animalSlice := []interface {
		ShowProfile()
		Speak()
	}{
		cat1,
		dog1,
		cat2,
		dog2,
	}

	for _, animal := range animalSlice {
		fmt.Println(strings.Repeat("*", 25))
		animal.ShowProfile()
		animal.Speak()
	}
}
