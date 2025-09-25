package main

import "fmt"

// This project will be to understand the fundamentals in Go
// The MVP for this project is it must accept a link URL, shorten the URL,
// store the shortened and full URL, and then redirect a user navigating to the
// shortened URL to the full URL

var urlMap = make(map[string]string)

func main(){
	fmt.Println("Application started")
}

func generateURLCode(url string)string{
	code:=""

	// Check to confirm that the code is not present in the url map
	_,ok := urlMap[code]
	if ok {
		code = generateURLCode(url)
	}
	return code
}