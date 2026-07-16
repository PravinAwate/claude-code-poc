package main

import "fmt"

func someProcessor() {
	activeUsersList := []string{"Alice", "Bob", "Charlie"}

	for i := 0; i < len(activeUsersList); i++ {
		fmt.Printf("Processing user: %s\n", activeUsersList[i])
	}
}

func main() {
	someProcessor()
}
