package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcde"
	for i := 0; i < len(s); i++ {
		// The type of slice of string is byte
		b := s[i]
		fmt.Printf("%s: %v\n", string(b), b)
	}
	fmt.Println(strings.Repeat("#", 25))

	for _, r := range s {
		fmt.Printf("%s: %v\n", string(r), r)
	}
}
