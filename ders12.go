package main

import "fmt"

func main() {
	fmt.Println("Hello ...")
	identityAsk()

	fmt.Println("See you soon ...")

	identityAsk()

}

func identityAsk() {
	fmt.Println("Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Age: ")
	var age int
	fmt.Scanln(&age)

	if age <= 18 {
		fmt.Println("Sorry", name, "unfortunately you are under 18 years of age.")
	} else {
		fmt.Println("Welcome,", name)
	}
}
