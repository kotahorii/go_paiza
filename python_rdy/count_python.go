package pythonrdy

import (
	"io/ioutil"
	"log"
	"regexp"
)

func CountPython(filename, word string) int {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	strList := regexp.
		MustCompile(word).
		FindAllString(string(b), -1)
	return len(strList)
}
