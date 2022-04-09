package selftaught

func BinarySearch(a_list []int, n int) bool {
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
