package helpers

import (
	"math/rand"
)

func GenerateUrlCode(urlCodeLength int) string {
	code := ""
	charsList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for range urlCodeLength {
		c := rand.Intn(len(charsList))
		code += string(charsList[c])
	}
	return code
}