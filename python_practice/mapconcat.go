package pythonpractice

import "strings"

func MapConcat(fn func(string) string, seq string, sep string) (result string) {
	strSlice := []string{}
	for _, s := range seq {
		strSlice = append(strSlice, fn(string(s)))
	}

	result = strings.Join(strSlice, sep)
	return
}
