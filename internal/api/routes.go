package api

import (
	"fmt"
	"net/http"
	"text/template"

	"cb-dev.com/link-shortener/internal/helpers"
)

const duplicatesAllowed = true

func HandleRootOrDefault(w http.ResponseWriter, r *http.Request, urlMap map[string]string) {
	path := r.URL.Path[1:]
	if target, ok := urlMap[path]; ok {
		http.Redirect(w, r, target, http.StatusFound)
	} else if r.URL.Path == "/" {
		HandleEmptyRootPage(w, http.StatusAccepted)
	} else {
		fmt.Printf("ERROR: Redirect not found: %v\n", path)
		HandleErrorResponse(w, "Redirect Not Found", http.StatusNotFound)
	}
}

func HandleAddRoute(w http.ResponseWriter, r *http.Request, urlMap map[string]string, urlCodeLength int) {

	if r.Method != "POST" {
		HandleErrorResponse(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		HandleErrorResponse(w, "ERROR: Could Not Parse Form", http.StatusBadRequest)
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	form := r.MultipartForm
	values, ok := form.Value["url"]

	if !ok || len(values) == 0 {
		HandleErrorResponse(w, "Field 'url' missing", http.StatusBadRequest)
		return
	}

	desiredUrl := values[0]

	_, err := helpers.IsValidUrl(desiredUrl)
	if err != nil {
		message := "The URL \"" + desiredUrl + "\" doesn't look like a valid URL"
		HandleErrorResponse(w, message, http.StatusBadRequest)
		return
	}

	duplicateValue := helpers.ContainsValue(urlMap, desiredUrl)
	fmt.Printf("The URL is: %v\n", desiredUrl)
	fmt.Printf("The map contains URL '%v': %v\n", desiredUrl, duplicateValue)

	if duplicateValue && !duplicatesAllowed {
		HandleErrorResponse(w, "URL Already Exists", http.StatusConflict)
		return
	}

	// Create the short code for redirection and add it to the "database"
	shortCode := helpers.GenerateUrlCode(urlCodeLength)
	urlMap[shortCode] = desiredUrl

	// Extract the scheme and host from the request
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host

	// Create the full URL string to be presented to the user
	shortUrl := fmt.Sprintf("%s://%s/%s", scheme, host, shortCode)

	HandleGeneratedURL(w, shortUrl, http.StatusAccepted)
}

func HandleEmptyRootPage(w http.ResponseWriter, status int) {
	template, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		RedirectUrlHeader string
		RedirectUrl       string
	}{
		RedirectUrlHeader: "",
		RedirectUrl:       "",
	}

	w.WriteHeader(status)
	template.Execute(w, data)
}

func HandleGeneratedURL(w http.ResponseWriter, redirectUrl string, status int) {
	template, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		RedirectUrlHeader string
		RedirectUrl       string
	}{
		RedirectUrlHeader: "Your redirection URL:",
		RedirectUrl:       redirectUrl,
	}

	w.WriteHeader(status)
	template.Execute(w, data)
}

func HandleErrorResponse(w http.ResponseWriter, message string, status int) {
	template, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error()+"\n"+message, status)
		return
	}

	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: message,
	}

	w.WriteHeader(status)
	template.Execute(w, data)
}
