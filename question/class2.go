package question

import "fmt"

type Animal interface {
	ShowProfile()
	Speak()
}

func PrintAnimalName(a Animal) {
	a.ShowProfile()
}

func AnimalSpeak(a Animal) {
	a.Speak()
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

func (cat *Cat) Speak() {
	fmt.Println("nya~")
}

func (cat *Cat) Sleep() {
	fmt.Println("zzz...")
}

type Dog2 struct {
	name string
	age  int
}

func NewDog2(name string, age int) *Dog2 {
	return &Dog2{name, age}
}

func (dog *Dog2) ShowProfile() {
	fmt.Printf("%s's age is %d\n", dog.name, dog.age)
}

func (dog *Dog2) Speak() {
	fmt.Println("bow")
}

func (dog *Dog2) Run() {
	fmt.Println("Go straight")
}
