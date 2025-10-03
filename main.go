package main

import (
	"fmt"
	"os"

	helper "cb-dev.com/link-shortener/internal"
	"github.com/davecgh/go-spew/spew"
)

// This project will be to understand the fundamentals in Go
// The MVP for this project is it must accept a link URL, shorten the URL,
// store the shortened and full URL, and then redirect a user navigating to the
// shortened URL to the full URL

const UrlCodeLength = 8
const databaseFile = "database.json"

var urlMap = make(map[string]string)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	_, err := os.Stat(databaseFile)
	if !os.IsNotExist(err){
		helper.LoadData(databaseFile,urlMap)
	}

	temp := helper.GenerateUrlCode(UrlCodeLength)
	urlMap[temp] = os.Args[1]
	helper.SaveData(databaseFile,urlMap)
	
	fmt.Printf("The code is: %v\n", temp)
	spew.Dump(urlMap)
}

