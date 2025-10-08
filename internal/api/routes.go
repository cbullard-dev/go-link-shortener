package api

import (
	"fmt"
	"net/http"
)

func HandleRootOrDefault(w http.ResponseWriter, r *http.Request, urlMap map[string]string) {
	path := r.URL.Path[1:]
	if target, ok := urlMap[path]; ok {
		http.Redirect(w, r, target, http.StatusFound)
	} else if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		fmt.Printf(`ERROR: Redirect not found: %v`, path)
		http.Error(w, "Redirect Not Found", http.StatusNotFound)
	}
}
