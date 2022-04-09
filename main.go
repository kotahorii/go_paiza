package main

import (
	"fmt"
)

func main() {
	// fmt.Println(binary_search([]int{1, 2, 3, 4, 5}, 1))
	// fmt.Println(insertionSort([]int{5, 3, 1, 6, 2}))
	fmt.Println(mergeSort([]int{5, 3, 1, 6, 2}))
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

func insertionSort(a_list []int) []int {
	for i := 1; i < len(a_list); i++ {
		value := a_list[i]
		fmt.Println(value)
		for i > 0 && a_list[i-1] > value {
			a_list[i] = a_list[i-1]
			i -= 1
		}
		a_list[i] = value
	}
	return a_list
}

func mergeSort(a_list []int) []int {
	if len(a_list) > 1 {
		copied := make([]int, len(a_list))
		copy(copied, a_list)
		mid := len(a_list) / 2

		left_half, right_half := copied[:mid], copied[mid:]
		mergeSort(left_half)
		mergeSort(right_half)

		left_ind, right_ind, list_ind := 0, 0, 0
		for left_ind < len(left_half) && right_ind < len(right_half) {
			if left_half[left_ind] >= right_half[right_ind] {
				a_list[list_ind] = right_half[right_ind]
				right_ind += 1
			} else {
				a_list[list_ind] = left_half[left_ind]
				left_ind += 1
			}
			list_ind += 1
		}

		for left_ind < len(left_half) {
			a_list[list_ind] = left_half[left_ind]
			left_ind += 1
			list_ind += 1
		}

		for right_ind < len(right_half) {
			a_list[list_ind] = right_half[right_ind]
			right_ind += 1
			list_ind += 1
		}
	}
	return a_list
}
