package main

import "fmt"

func main() {

	/*var x int = 9

	if x > 9 {
		fmt.Println("x sayısı 9 dan büyüktür.")
	} else if x == 9 {
		fmt.Println("x sayısı 9 a eşittir")
	} else {
		fmt.Println("x sayısı 9 dan büyük değildir.")
	}*/

	ortalama := 97

	if ortalama > 90 {
		fmt.Println("A")
	} else if ortalama > 80 {
		fmt.Println("B")
	} else if ortalama > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
