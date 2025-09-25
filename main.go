package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// This project will be to understand the fundamentals in Go
// The MVP for this project is it must accept a link URL, shorten the URL,
// store the shortened and full URL, and then redirect a user navigating to the
// shortened URL to the full URL

var urlMap = make(map[string]string)
const urlCodeLength = 8

func main(){
	temp := generateURLCode()
	urlMap[temp] = os.Args[1]
	fmt.Printf("The code is: %v\n", temp)
	spew.Dump(urlMap)
}

func generateURLCode()string{
	code:=""
	charsList:="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for range(urlCodeLength){
		c, _ := rand.Int(rand.Reader,big.NewInt(int64(len(charsList))))
		code += string(charsList[c.Int64()])
	}

	// Check to confirm that the code is not present in the url map
	_,ok := urlMap[code]
	if ok {
		code = generateURLCode()
	}
	return code
}