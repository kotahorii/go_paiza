package question

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountNumberOfTimes() {
	sc := bufio.NewScanner(os.Stdin)

	if sc.Scan() {
		strs := sc.Text()
		split_strs := strings.Split(strs, " ")
		fmt.Println(split_strs)

		strs_count := make(map[string]int)

		for _, str := range split_strs {
			if _, ok := strs_count[str]; !ok {
				strs_count[str] = 1
			} else {
				strs_count[str] += 1
			}
		}
		printCounts(strs_count)
	}
}

func printCounts(strs_count map[string]int) {
	for k, v := range strs_count {
		fmt.Printf("%s, %d\n", k, v)
	}
}
