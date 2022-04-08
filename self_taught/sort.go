package selftaught

func InsertionSort(a_list []int) []int {
	for i := 1; i < len(a_list); i++ {
		value := a_list[i]

		for i > 0 && a_list[i-1] > value {
			a_list[i] = a_list[i-1]
			i = i - 1
		}
		a_list[i] = value
	}
	return a_list
}

func MergeSort(a_list []int) []int {
	if len(a_list) > 1 {
		mid := len(a_list) / 2
		copied_list := make([]int, len(a_list))
		// use copy function to copy the slice
		// NG!! copied_list := a_list
		copy(copied_list, a_list)
		left_half := copied_list[:mid]
		right_half := copied_list[mid:]

		MergeSort(left_half)
		MergeSort(right_half)

		left_index := 0
		right_index := 0
		a_list_index := 0

		for left_index < len(left_half) && right_index < len(right_half) {
			if left_half[left_index] <= right_half[right_index] {
				a_list[a_list_index] = left_half[left_index]
				left_index += 1
			} else {
				a_list[a_list_index] = right_half[right_index]
				right_index += 1
			}

			a_list_index += 1
		}

		for left_index < len(left_half) {
			a_list[a_list_index] = left_half[left_index]
			left_index += 1
			a_list_index += 1
		}

		for right_index < len(right_half) {
			a_list[a_list_index] = right_half[right_index]
			right_index += 1
			a_list_index += 1
		}
	}
	return a_list
}
