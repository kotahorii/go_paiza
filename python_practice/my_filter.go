package pythonpractice

func IsEven(n int) bool {
	return n%2 == 0
}

func IsScoreGreaterThan60(n int) bool {
	return n >= 60
}

func MyFilter(f func(int) bool, list []int) (result []int) {
	for _, n := range list {
		if f(n) {
			result = append(result, n)
		}
	}
	return
}

func Partition(f func(int) bool, list []int) (result [][]int) {
	var trueList, falseList []int
	for _, n := range list {
		if f(n) {
			trueList = append(trueList, n)
		} else {
			falseList = append(falseList, n)
		}
	}
	result = append(result, trueList, falseList)
	return
}

func PartitionScore(f func(int) bool, scoreAndName map[string]int) (result []map[string]int) {
	trueList, falseList := map[string]int{}, map[string]int{}

	for k, v := range scoreAndName {
		if f(v) {
			trueList[k] = v
		} else {
			falseList[k] = v
		}
	}
	result = append(result, trueList, falseList)
	return
}
