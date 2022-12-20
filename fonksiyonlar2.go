package main

import "fmt"

func main() {
	fmt.Println("Sayı Giriniz: ")
	var y int
	fmt.Scanln(&y)
	fmt.Println(kareAl(y))

	// Burdan itibaren toplama işlemi

	fmt.Println("Sayı giriniz: ")
	var number1 int
	fmt.Scanln(&number1)

	fmt.Println("Sayı giriniz: ")
	var number2 int
	fmt.Scanln(&number2)

	fmt.Println(topla(number1, number2))
}

func kareAl(x int) int {
	return x * x
}

func topla(number1, number2 int) int {
	fmt.Println("Toplam: ")
	return number1 + number2
}
