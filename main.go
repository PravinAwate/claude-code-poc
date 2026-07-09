package main

import (
	"fmt"
	"net/http"
)

// Demo Secret (Pipeline should flag this)
const ApiKeyPlaceholder = "sk_live_1234567890abcdef" 

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Code review should flag this: No error handling or proper header setting
		fmt.Fprintf(w, "Hello, World!") 
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
