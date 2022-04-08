package selftaught

import (
	"fmt"
	"strings"
)

func IsPalindrome(word string) bool {
	return reverse(strings.ToLower(word)) == word
}

func reverse(s string) string {
	runes := []rune(s)
	fmt.Println(runes)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
