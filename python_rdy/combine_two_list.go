package pythonrdy

func CombineTwoList(l1, l2 []string) (result [][]string) {
	for _, v1 := range l1 {
		for _, v2 := range l2 {
			result = append(result, []string{v1, v2})
		}
	}
	return
}

func MakeNEmptyList(n int) (result [][]interface{}) {
	for i := 0; i < n; i++ {
		result = append(result, make([]interface{}, 0))
	}
	return
}
