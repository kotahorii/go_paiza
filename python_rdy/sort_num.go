package pythonrdy

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func SortNumInTextFile(filename string) (result []int) {
	// read file
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// split string
	numList := regexp.
		MustCompile("[0-9]+").
		FindAllString(string(b), -1)

	for _, str := range numList {
		if n, err := strconv.Atoi(str); err == nil {
			result = append(result, n)
		}
	}

	// sort by decending
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return
}
