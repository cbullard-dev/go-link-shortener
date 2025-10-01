package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// This project will be to understand the fundamentals in Go
// The MVP for this project is it must accept a link URL, shorten the URL,
// store the shortened and full URL, and then redirect a user navigating to the
// shortened URL to the full URL

const UrlCodeLength = 8

var urlMap = make(map[string]string)

func main() {
	if len(os.Args)<=1{
		return
	}
	temp := generateUrlCode()
	urlMap[temp] = os.Args[1]
	fmt.Printf("The code is: %v\n", temp)
	spew.Dump(urlMap)
}

func generateUrlCode() string {
	code := ""
	charsList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for range UrlCodeLength {
		c := rand.Intn(len(charsList))
		code += string(charsList[c])
	}

	// Check to confirm that the code is not present in the url map
	_, ok := urlMap[code]
	if ok {
		code = generateUrlCode()
	}
	return code
}
