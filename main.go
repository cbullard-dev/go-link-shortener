package main

import (
	"log"
	"net/http"
	"os"

	"cb-dev.com/link-shortener/internal/api"
	storage "cb-dev.com/link-shortener/internal/storage"
)

// This project will be to understand the fundamentals in Go
// The MVP for this project is it must accept a link URL, shorten the URL,
// store the shortened and full URL, and then redirect a user navigating to the
// shortened URL to the full URL

const UrlCodeLength = 8
const databaseFile = "database.json"

var urlMap = make(map[string]string)

// For the time being I'm going to use the built-in default mux
// I might add my own mux later
// var mux = http.NewServeMux()

func main() {

	_, err := os.Stat(databaseFile)
	if !os.IsNotExist(err) {
		storage.LoadData(databaseFile, urlMap)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.HandleRootOrDefault(w, r, urlMap)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		api.HandleAddRoute(w, r, urlMap, UrlCodeLength)
		storage.SaveData(databaseFile, urlMap)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
