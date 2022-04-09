package main

import "fmt"

func main() {
	fmt.Println(binary_search([]int{1, 2, 3, 4, 5}, 1))
}

func binary_search(a_list []int, n int) bool {
	first, last := 0, len(a_list)-1

	for last >= first {
		mid := (first + last) / 2
		if n == a_list[mid] {
			return true
		}

		if n > a_list[mid] {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return false
}
