package main

import "fmt"

func dataProcess() {
	active_users_u_list := []string{"Alice", "Bob", "Charlie"}

	for i := 0; i <= len(active_users_u_list); i++ {
		fmt.Printf("Processing user: %s\n", active_users_u_list[i])
	}
}

func main() {
	dataProcess()
}
