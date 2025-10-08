package api

import (
	"fmt"
	"net/http"

	"cb-dev.com/link-shortener/internal/helpers"
)

const duplicatesAllowed = true

func HandleRootOrDefault(w http.ResponseWriter, r *http.Request, urlMap map[string]string) {
	path := r.URL.Path[1:]
	if target, ok := urlMap[path]; ok {
		http.Redirect(w, r, target, http.StatusFound)
	} else if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		fmt.Printf("ERROR: Redirect not found: %v\n", path)
		http.Error(w, "Redirect Not Found", http.StatusNotFound)
	}
}

func HandleAddRoute(w http.ResponseWriter, r *http.Request, urlMap map[string]string, urlCodeLength int) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "ERROR: Could Not Parse Form", http.StatusBadRequest)
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	form := r.MultipartForm
	values, ok := form.Value["url"]

	if !ok || len(values) == 0 {
		http.Error(w, "Field 'url' missing", http.StatusBadRequest)
		return
	}

	desiredUrl := values[0]
	duplicateValue := helpers.ContainsValue(urlMap, desiredUrl)
	fmt.Printf("The URL is: %v\n", desiredUrl)
	fmt.Printf("The map contains URL '%v': %v\n", desiredUrl, duplicateValue)

	if duplicateValue && !duplicatesAllowed {
		http.Error(w, "URL Already Exists", http.StatusConflict)
		return
	}

	temp := helpers.GenerateUrlCode(urlCodeLength)
	urlMap[temp] = desiredUrl
	http.Redirect(w, r, "/", http.StatusFound)
}
