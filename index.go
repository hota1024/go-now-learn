package handler

import (
	"fmt"
	"net/http"
)

// Handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>index.go Only!</h1>")
}
