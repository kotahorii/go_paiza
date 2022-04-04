package pythonpractice

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func RenameFile(oldExt, newExt string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == oldExt {
			fmt.Printf(
				"%s => %s%s",
				file.Name(),
				strings.Replace(file.Name(), oldExt, "", -1),
				newExt,
			)
		}
	}
}
