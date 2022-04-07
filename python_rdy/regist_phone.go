package pythonrdy

import (
	"bufio"
	"fmt"
	"os"
)

func Reg() map[string]string {
	phones := make(map[string]string, 0)
	defer fmt.Println(phones)

	sc := bufio.NewScanner(os.Stdin)
	var name, tel string

	for {
		fmt.Print("name: ")
		name = scanText(sc)

		fmt.Print("tel: ")
		tel = scanText(sc)

		phones[name] = tel
	}
}

func scanText(sc *bufio.Scanner) string {
	if sc.Scan() {
		if sc.Text() == "phone" {
			panic("")
		}
		return sc.Text()
	}
	return ""
}
