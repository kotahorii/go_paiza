package question

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func TwoTimes(nums []int) []int {
	for i := range nums {
		nums[i] *= 2
	}
	return nums
}

func ReverseNums(nums []int) []int {
	max_len := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[max_len-i] = nums[max_len-i], nums[i]
	}
	return nums
}

func DivideEvenOdd(nums []int) (evenNums, oddNums []int) {
	for i := range nums {
		if nums[i]%2 == 0 {
			evenNums = append(evenNums, nums[i])
		} else {
			oddNums = append(oddNums, nums[i])
		}
	}

	return
}

func Return10Nums() (result []int) {
	sc := bufio.NewScanner(os.Stdin)
	sum := 0
	for i := 0; i < 10; i++ {
		if sum >= 100 {
			break
		}

		if sc.Scan() {
			num, err := strconv.Atoi(sc.Text())
			if err != nil {
				log.Fatal(err)
			}

			result = append(result, num)
			sum += num
		}
	}
	return
}

func TimesTable() (table [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return
}

func CalcProduct(num1, num2 int) int {
	kuku := TimesTable()
	return kuku[num1-1][num2-1]
}
