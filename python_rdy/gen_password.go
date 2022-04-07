package pythonrdy

import "math/rand"

func GenPassword(lenPassword int) (password string) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for len(password) < lenPassword {
		password += string(chars[rand.Intn(len(chars)-1)])
	}
	return
}
