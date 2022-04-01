package question

import "fmt"

type Animal interface {
	ShowProfile()
}

type Cat struct {
	name string
	age  int
}

func NewCat(name string, age int) *Cat {
	return &Cat{name, age}
}

func (cat *Cat) ShowProfile() {
	fmt.Printf("%s's age is %d\n", cat.name, cat.age)
}

func (cat *Cat) Sleep() {
	fmt.Println("zzz...")
}
