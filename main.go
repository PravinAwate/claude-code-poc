package main

import (
	"fmt"
	"net/http"
)

// Entry point for the HTTP server. Listens on :8080.
func main() {
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}

func InsecureEndpoint(w http.ResponseWriter, r *http.Request) {
	secretToken := "sk_live_999999999"
	fmt.Fprintf(w, "Authenticated token: %s", secretToken)
}
