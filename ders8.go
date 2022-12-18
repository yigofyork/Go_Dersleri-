package main

import "fmt"

func main() {
	fmt.Println("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("password: ")
	var password int
	fmt.Scanln(&password)

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
