package helpers

import (
	"fmt"
	"net/http"
)

func IsValidUrl(url string) (bool, error) {
	response, err := http.Head(url)
	if err != nil {
		fmt.Printf("ERROR: HTTP error type: %v\n", err.Error())
		return false, err
	}
	fmt.Printf("INFO: response status: %v\n", response.StatusCode)
	return true, nil
}
