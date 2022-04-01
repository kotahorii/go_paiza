package question

import (
	"errors"
	"fmt"
)

type Dog struct {
	name  string
	age   int
	breed string
}

type MyError interface {
	Error()
}

func NewDog(name string, age int, breed string) *Dog {
	return &Dog{name: name, age: age, breed: breed}
}

func (dog *Dog) ShowProfile() {
	fmt.Printf(
		"%s's age is %d. Breed is %s",
		dog.name, dog.age, dog.breed,
	)
}

type CoinCase struct {
	one, five, ten, fifty, hundred, fiveHundred int
}

func NewCoinCase() *CoinCase {
	return &CoinCase{0, 0, 0, 0, 0, 0}
}

func (coinCase *CoinCase) AddCoins(kind, number int) error {
	switch kind {
	case 1:
		coinCase.one += number
	case 5:
		coinCase.five += number
	case 10:
		coinCase.ten += number
	case 50:
		coinCase.fifty += number
	case 100:
		coinCase.hundred += number
	case 500:
		coinCase.fiveHundred += number
	default:
		return errors.New("please input valid kind")
	}
	return nil
}

func (coinCase *CoinCase) GetCounts(kind int) (number int, err error) {
	switch kind {
	case 1:
		return coinCase.one, nil
	case 5:
		return coinCase.five, nil
	case 10:
		return coinCase.ten, nil
	case 50:
		return coinCase.fifty, nil
	case 100:
		return coinCase.hundred, nil
	case 500:
		return coinCase.fiveHundred, nil
	default:
		return 0, errors.New("please input valid type of money")
	}
}
