package pythonrdy

import (
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CalcSin(degree int) float64 {
	return math.Round(math.Sin(float64(degree)*math.Pi/180)*10) / 10
}

func SortNumInTextFile(filename string) (result []int) {
	// read file
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// split string
	for _, str := range strings.Split(string(b), " ") {
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
