package utils

import (
	"net/http"
)

// Check if the user has an internet connection
func HasInternetConnection() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}
