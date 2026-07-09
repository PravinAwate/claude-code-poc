package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
