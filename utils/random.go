package utils

import "math/rand"

var runes = []rune("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func RandomURL(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}
