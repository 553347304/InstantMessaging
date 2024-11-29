package randoms

import (
	"math/rand"
)

var char = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func String(length int) string {
	num := make([]rune, length)
	for i := range num {
		num[i] = char[rand.Intn(len(char))]
	}
	return string(num)
}
