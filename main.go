package main

import "fmt"

func main() {
	users := []string{"Alice", "Bob", "Charlie"}

	// 🛑 CRITICAL BUG: The condition "i <= len(users)" causes an out-of-bounds crash.
	// In Go, slices are 0-indexed. The last valid element is at index 2 (len - 1).
	// When i equals 3, the program will panic and crash.
	for i := 0; i <= len(users); i++ {
		fmt.Printf("Processing user: %s\n", users[i])
	}
}
