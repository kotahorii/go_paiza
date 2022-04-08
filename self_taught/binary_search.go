package selftaught

func BinarySearch(a_list []int, n int) bool {
	first := 0
	last := len(a_list) - 1

	for last >= first {
		mid := (first + last) / 2
		if n == a_list[mid] {
			return true
		}

		if n > a_list[mid] {
			first = a_list[mid] + 1
		} else {
			last = a_list[mid] - 1
		}
	}
	return false
}
